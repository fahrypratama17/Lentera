package service

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/contract"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/entity"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

type materialService struct {
	repo     contract.MaterialRepository
	uploader contract.FileUploader
	rt       contract.Broadcaster
	now      func() time.Time
}

func New(repo contract.MaterialRepository, rt contract.Broadcaster, now func() time.Time) contract.MaterialService {
	return NewWithUploader(repo, nil, rt, now)
}

func NewWithUploader(repo contract.MaterialRepository, uploader contract.FileUploader, rt contract.Broadcaster, now func() time.Time) contract.MaterialService {
	if rt == nil {
		rt = contract.NoopBroadcaster{}
	}
	if now == nil {
		now = time.Now
	}
	return &materialService{
		repo:     repo,
		uploader: uploader,
		rt:       rt,
		now:      now,
	}
}

func (s *materialService) ListMaterials(ctx context.Context, userID, cohortID uint, q dto.ListMaterialsQuery) ([]dto.MaterialResponse, error) {
	role, err := s.repo.GetActiveRole(ctx, cohortID, userID)
	if err != nil {
		return nil, err
	}

	materials, err := s.repo.ListMaterials(ctx, cohortID, strings.TrimSpace(q.Q))
	if err != nil {
		return nil, err
	}

	now := s.now()
	out := make([]dto.MaterialResponse, 0, len(materials))

	// Mentor: tanpa progres pribadi, filter status/late tidak berlaku
	if role != entity.RoleMember {
		for i := range materials {
			out = append(out, toMaterialResponse(&materials[i], nil, now))
		}
		return out, nil
	}

	progresses, err := s.repo.ListUserProgress(ctx, cohortID, userID)
	if err != nil {
		return nil, err
	}
	byMaterial := make(map[uint]*entity.Progress, len(progresses))
	for i := range progresses {
		byMaterial[progresses[i].MaterialID] = &progresses[i]
	}

	names, err := s.reviewerNames(ctx, progresses)
	if err != nil {
		return nil, err
	}

	for i := range materials {
		m := &materials[i]
		p := toProgressResponse(m, byMaterial[m.ID], userID, now, names)

		if q.Status != "" && p.Status != q.Status {
			continue
		}
		if q.Late != nil && p.IsLate != *q.Late {
			continue
		}
		out = append(out, toMaterialResponse(m, &p, now))
	}
	return out, nil
}

func (s *materialService) GetMaterial(ctx context.Context, userID, cohortID, materialID uint) (*dto.MaterialResponse, error) {
	role, err := s.repo.GetActiveRole(ctx, cohortID, userID)
	if err != nil {
		return nil, err
	}

	m, err := s.repo.GetMaterial(ctx, cohortID, materialID)
	if err != nil {
		return nil, err
	}

	now := s.now()
	if role != entity.RoleMember {
		res := toMaterialResponse(m, nil, now)
		return &res, nil
	}

	p, err := s.repo.GetUserProgress(ctx, materialID, userID)
	if err != nil {
		return nil, err
	}

	var list []entity.Progress
	if p != nil {
		list = []entity.Progress{*p}
	}
	names, err := s.reviewerNames(ctx, list)
	if err != nil {
		return nil, err
	}

	pr := toProgressResponse(m, p, userID, now, names)
	res := toMaterialResponse(m, &pr, now)
	return &res, nil
}

func (s *materialService) MySummary(ctx context.Context, userID, cohortID uint) (*dto.PersonalSummary, error) {
	role, err := s.repo.GetActiveRole(ctx, cohortID, userID)
	if err != nil {
		return nil, err
	}
	if role != entity.RoleMember {
		return nil, response.ErrForbidden
	}

	materials, err := s.repo.ListMaterials(ctx, cohortID, "")
	if err != nil {
		return nil, err
	}
	progresses, err := s.repo.ListUserProgress(ctx, cohortID, userID)
	if err != nil {
		return nil, err
	}
	byMaterial := make(map[uint]*entity.Progress, len(progresses))
	for i := range progresses {
		byMaterial[progresses[i].MaterialID] = &progresses[i]
	}

	now := s.now()
	sum := &dto.PersonalSummary{TotalMaterials: len(materials)}

	for i := range materials {
		m := &materials[i]
		status := entity.StatusNotStarted
		if p := byMaterial[m.ID]; p != nil {
			status = p.Status
		}

		switch status {
		case entity.StatusNotStarted:
			sum.NotStarted++
		case entity.StatusInProgress:
			sum.InProgress++
		case entity.StatusSubmitted:
			sum.Submitted++
		case entity.StatusCompleted:
			sum.Completed++
		}

		if status != entity.StatusCompleted && isPastDeadline(m, now) {
			sum.Late++
		}
	}

	if sum.TotalMaterials > 0 {
		sum.CompletionRate = float64(sum.Completed) / float64(sum.TotalMaterials)
	}
	return sum, nil
}

// kumpulkan ID reviewer unik lalu ambil namanya sekali query
func (s *materialService) reviewerNames(ctx context.Context, progresses []entity.Progress) (map[uint]string, error) {
	seen := map[uint]struct{}{}
	ids := make([]uint, 0)
	for _, p := range progresses {
		if p.ReviewerID == nil {
			continue
		}
		if _, ok := seen[*p.ReviewerID]; ok {
			continue
		}
		seen[*p.ReviewerID] = struct{}{}
		ids = append(ids, *p.ReviewerID)
	}
	return s.repo.GetUserNames(ctx, ids)
}

// ───────────────── helper ─────────────────

func cohortRoom(cohortID uint) string { return fmt.Sprintf("cohort:%d", cohortID) }
func mentorRoom(cohortID uint) string { return fmt.Sprintf("cohort:%d:mentors", cohortID) }

func (s *materialService) requireMentor(ctx context.Context, cohortID, userID uint) error {
	role, err := s.repo.GetActiveRole(ctx, cohortID, userID)
	if err != nil {
		return err
	}
	if role != entity.RoleMentor {
		return response.ErrForbidden
	}
	return nil
}

// event dashboard cukup sinyal; data dashboard diambil ulang lewat REST
func (s *materialService) notifyDashboard(cohortID uint) {
	s.rt.Broadcast(mentorRoom(cohortID), "dashboard.updated", map[string]any{"cohort_id": cohortID})
}

// ───────────────── CRUD materi (Mentor) ─────────────────

func (s *materialService) CreateMaterial(ctx context.Context, userID, cohortID uint, req dto.CreateMaterialRequest) (*dto.MaterialResponse, error) {
	if err := s.requireMentor(ctx, cohortID, userID); err != nil {
		return nil, err
	}

	title := strings.TrimSpace(req.Title)
	guide := strings.TrimSpace(req.Guide)
	if len(title) < 3 || guide == "" {
		return nil, fmt.Errorf("%w: title minimal 3 karakter dan guide wajib diisi", response.ErrBadRequest)
	}

	pos, err := s.repo.NextPosition(ctx, cohortID)
	if err != nil {
		return nil, err
	}

	refs := req.References
	if refs == nil {
		refs = []string{}
	}

	m := &entity.Material{
		CohortID:   cohortID,
		Title:      title,
		Guide:      guide,
		References: refs,
		Deadline:   toUTC(req.Deadline),
		Position:   pos,
	}
	if err := s.repo.CreateMaterial(ctx, m); err != nil {
		return nil, err
	}

	slog.InfoContext(ctx, "material created",
		"actor_id", userID, "cohort_id", cohortID, "material_id", m.ID, "position", m.Position)

	res := toMaterialResponse(m, nil, s.now())
	s.rt.Broadcast(cohortRoom(cohortID), "material.created", res)
	s.notifyDashboard(cohortID)
	return &res, nil
}

func (s *materialService) UpdateMaterial(ctx context.Context, userID, cohortID, materialID uint, req dto.UpdateMaterialRequest) (*dto.MaterialResponse, error) {
	if err := s.requireMentor(ctx, cohortID, userID); err != nil {
		return nil, err
	}

	fields := map[string]any{}

	if req.Title != nil {
		t := strings.TrimSpace(*req.Title)
		if len(t) < 3 {
			return nil, fmt.Errorf("%w: title minimal 3 karakter", response.ErrBadRequest)
		}
		fields["title"] = t
	}
	if req.Guide != nil {
		g := strings.TrimSpace(*req.Guide)
		if g == "" {
			return nil, fmt.Errorf("%w: guide tidak boleh kosong", response.ErrBadRequest)
		}
		fields["guide"] = g
	}
	if req.References != nil {
		fields["references"] = req.References
	}
	if req.Deadline.Set {
		if req.Deadline.Value == nil {
			fields["deadline"] = nil // hapus deadline
		} else {
			fields["deadline"] = req.Deadline.Value.UTC()
		}
	}

	if len(fields) == 0 {
		return nil, fmt.Errorf("%w: tidak ada field yang diubah", response.ErrBadRequest)
	}

	if err := s.repo.UpdateMaterial(ctx, cohortID, materialID, fields); err != nil {
		return nil, err
	}

	changed := make([]string, 0, len(fields))
	for k := range fields {
		changed = append(changed, k)
	}
	slog.InfoContext(ctx, "material updated",
		"actor_id", userID, "cohort_id", cohortID, "material_id", materialID, "fields", changed)

	m, err := s.repo.GetMaterial(ctx, cohortID, materialID)
	if err != nil {
		return nil, err
	}

	res := toMaterialResponse(m, nil, s.now())
	s.rt.Broadcast(cohortRoom(cohortID), "material.updated", res)
	s.notifyDashboard(cohortID) // deadline berubah → jumlah terlambat bisa berubah
	return &res, nil
}

func (s *materialService) DeleteMaterial(ctx context.Context, userID, cohortID, materialID uint) error {
	if err := s.requireMentor(ctx, cohortID, userID); err != nil {
		return err
	}

	if err := s.repo.DeleteMaterial(ctx, cohortID, materialID); err != nil {
		return err
	}

	slog.InfoContext(ctx, "material deleted",
		"actor_id", userID, "cohort_id", cohortID, "material_id", materialID)

	s.rt.Broadcast(cohortRoom(cohortID), "material.deleted", map[string]any{"id": materialID})
	s.notifyDashboard(cohortID)
	return nil
}

func (s *materialService) ReorderMaterials(ctx context.Context, userID, cohortID uint, req dto.ReorderMaterialsRequest) ([]dto.MaterialResponse, error) {
	if err := s.requireMentor(ctx, cohortID, userID); err != nil {
		return nil, err
	}

	if err := s.repo.ReorderMaterials(ctx, cohortID, req.MaterialIDs); err != nil {
		return nil, err
	}

	slog.InfoContext(ctx, "materials reordered",
		"actor_id", userID, "cohort_id", cohortID, "count", len(req.MaterialIDs))

	materials, err := s.repo.ListMaterials(ctx, cohortID, "")
	if err != nil {
		return nil, err
	}

	now := s.now()
	out := make([]dto.MaterialResponse, 0, len(materials))
	for i := range materials {
		out = append(out, toMaterialResponse(&materials[i], nil, now))
	}

	s.rt.Broadcast(cohortRoom(cohortID), "material.reordered", map[string]any{"material_ids": req.MaterialIDs})
	return out, nil
}

func toUTC(t *time.Time) *time.Time {
	if t == nil {
		return nil
	}
	u := t.UTC()
	return &u
}
