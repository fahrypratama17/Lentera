package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/contract"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/entity"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

type reviewService struct {
	repo contract.ReviewRepository
	rt   contract.Broadcaster
	now  func() time.Time
}

func NewReview(repo contract.ReviewRepository, rt contract.Broadcaster, now func() time.Time) contract.ReviewService {
	if rt == nil {
		rt = contract.NoopBroadcaster{}
	}
	if now == nil {
		now = time.Now
	}
	return &reviewService{repo: repo, rt: rt, now: now}
}

// ───────────────── helper ─────────────────

func (s *reviewService) requireMentor(ctx context.Context, cohortID, userID uint) error {
	role, err := s.repo.GetActiveRole(ctx, cohortID, userID)
	if err != nil {
		return err
	}
	if role != entity.RoleMentor {
		return response.ErrForbidden
	}
	return nil
}

func userRoom(userID uint) string { return fmt.Sprintf("user:%d", userID) }

func toReviewItem(m *entity.Material, p *entity.Progress, mem entity.MemberInfo, now time.Time, names map[uint]string) dto.ReviewItemResponse {
	return dto.ReviewItemResponse{
		ProgressResponse: toProgressResponse(m, p, p.UserID, now, names),
		Member:           dto.ReviewMember{UserID: mem.UserID, Name: mem.Name, Status: mem.Status},
		Material:         toMatrixMaterial(m),
	}
}

// rakit satu ReviewItemResponse lengkap untuk satu progres
func (s *reviewService) buildItem(ctx context.Context, cohortID uint, p *entity.Progress) (*dto.ReviewItemResponse, error) {
	materials, err := s.repo.ListMaterials(ctx, cohortID, "")
	if err != nil {
		return nil, err
	}
	var m *entity.Material
	for i := range materials {
		if materials[i].ID == p.MaterialID {
			m = &materials[i]
			break
		}
	}
	if m == nil {
		return nil, response.ErrNotFound
	}

	members, err := s.repo.ListCohortMembers(ctx, cohortID)
	if err != nil {
		return nil, err
	}
	mem := entity.MemberInfo{UserID: p.UserID}
	for _, x := range members {
		if x.UserID == p.UserID {
			mem = x
			break
		}
	}

	names, err := s.repo.GetUserNames(ctx, reviewerIDsOf([]*entity.Progress{p}))
	if err != nil {
		return nil, err
	}

	item := toReviewItem(m, p, mem, s.now(), names)
	return &item, nil
}

// ───────────────── daftar & detail ─────────────────

func (s *reviewService) ListReviews(ctx context.Context, userID, cohortID uint, q dto.ListReviewsQuery) ([]dto.ReviewItemResponse, error) {
	if err := s.requireMentor(ctx, cohortID, userID); err != nil {
		return nil, err
	}

	status := q.Status
	if status == "" {
		status = entity.StatusSubmitted
	}

	progresses, err := s.repo.ListProgressByStatus(ctx, cohortID, status, q.MaterialID)
	if err != nil {
		return nil, err
	}
	materials, err := s.repo.ListMaterials(ctx, cohortID, "")
	if err != nil {
		return nil, err
	}
	members, err := s.repo.ListCohortMembers(ctx, cohortID)
	if err != nil {
		return nil, err
	}

	materialByID := make(map[uint]*entity.Material, len(materials))
	for i := range materials {
		materialByID[materials[i].ID] = &materials[i]
	}
	memberByID := make(map[uint]entity.MemberInfo, len(members))
	for _, mem := range members {
		memberByID[mem.UserID] = mem
	}

	ptrs := make([]*entity.Progress, 0, len(progresses))
	for i := range progresses {
		ptrs = append(ptrs, &progresses[i])
	}
	names, err := s.repo.GetUserNames(ctx, reviewerIDsOf(ptrs))
	if err != nil {
		return nil, err
	}

	search := strings.ToLower(strings.TrimSpace(q.Q))
	now := s.now()
	out := make([]dto.ReviewItemResponse, 0, len(progresses))

	for _, p := range ptrs {
		m, ok := materialByID[p.MaterialID]
		if !ok {
			continue
		}
		mem, ok := memberByID[p.UserID]
		if !ok {
			continue
		}
		if !q.IncludeInactive && mem.Status != "active" {
			continue
		}
		if search != "" && !strings.Contains(strings.ToLower(mem.Name), search) {
			continue
		}
		out = append(out, toReviewItem(m, p, mem, now, names))
	}
	return out, nil
}

func (s *reviewService) GetReviewItem(ctx context.Context, userID, cohortID, progressID uint) (*dto.ReviewItemResponse, error) {
	role, err := s.repo.GetActiveRole(ctx, cohortID, userID)
	if err != nil {
		return nil, err
	}

	p, err := s.repo.GetProgress(ctx, cohortID, progressID)
	if err != nil {
		return nil, err
	}
	// Mentor boleh lihat semua; member hanya progres miliknya
	if role != entity.RoleMentor && p.UserID != userID {
		return nil, response.ErrForbidden
	}
	return s.buildItem(ctx, cohortID, p)
}

// ───────────────── keputusan review ─────────────────

func (s *reviewService) Approve(ctx context.Context, userID, cohortID, progressID uint, req dto.ApproveRequest) (*dto.ReviewItemResponse, error) {
	var note *string
	if req.Note != nil {
		if n := strings.TrimSpace(*req.Note); n != "" {
			note = &n
		}
	}
	return s.decide(ctx, userID, cohortID, progressID, func(p *entity.Progress, now time.Time) error {
		return p.Approve(userID, note, now)
	})
}

func (s *reviewService) RequestRevision(ctx context.Context, userID, cohortID, progressID uint, req dto.ReviewNoteRequest) (*dto.ReviewItemResponse, error) {
	return s.decide(ctx, userID, cohortID, progressID, func(p *entity.Progress, now time.Time) error {
		return p.RequestRevision(userID, req.Note, now)
	})
}

func (s *reviewService) Reopen(ctx context.Context, userID, cohortID, progressID uint, req dto.ReviewNoteRequest) (*dto.ReviewItemResponse, error) {
	return s.decide(ctx, userID, cohortID, progressID, func(p *entity.Progress, now time.Time) error {
		return p.Reopen(userID, req.Note, now)
	})
}

// alur sama untuk semua keputusan: cek mentor → ambil progres → entity ubah status → simpan bersyarat → broadcast
func (s *reviewService) decide(
	ctx context.Context, userID, cohortID, progressID uint,
	apply func(p *entity.Progress, now time.Time) error,
) (*dto.ReviewItemResponse, error) {
	if err := s.requireMentor(ctx, cohortID, userID); err != nil {
		return nil, err
	}

	p, err := s.repo.GetProgress(ctx, cohortID, progressID)
	if err != nil {
		return nil, err
	}

	fromStatus := p.Status
	if err := apply(p, s.now()); err != nil {
		return nil, err // transisi tidak valid / catatan kosong — belum ada yang tersimpan
	}

	if err := s.repo.SaveReview(ctx, p, fromStatus); err != nil {
		return nil, err // didahului mentor lain → 409
	}

	item, err := s.buildItem(ctx, cohortID, p)
	if err != nil {
		return nil, err
	}

	// event cuma sinyal (tanpa catatan review) — detail diambil lewat REST yang sudah cek akses
	signal := map[string]any{
		"progress_id": p.ID,
		"material_id": p.MaterialID,
		"user_id":     p.UserID,
		"status":      p.Status,
	}
	s.rt.Broadcast(userRoom(p.UserID), "review.decided", signal)
	s.rt.Broadcast(mentorRoom(cohortID), "review.decided", signal)
	s.rt.Broadcast(mentorRoom(cohortID), "dashboard.updated", map[string]any{"cohort_id": cohortID})

	return item, nil
}
