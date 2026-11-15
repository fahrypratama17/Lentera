package service

import (
	"context"
	"strings"
	"time"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/contract"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/entity"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

type monitoringService struct {
	repo contract.MonitoringRepository
	now  func() time.Time
}

func NewMonitoring(repo contract.MonitoringRepository, now func() time.Time) contract.MonitoringService {
	if now == nil {
		now = time.Now
	}
	return &monitoringService{repo: repo, now: now}
}

// ───────────────── helper ─────────────────

type progressKey struct{ userID, materialID uint }

// snapshot = semua data yang dibutuhkan matriks & dashboard, diambil dengan 3 query
type cohortSnapshot struct {
	materials []entity.Material
	members   []entity.MemberInfo
	progress  map[progressKey]*entity.Progress
}

func (s *monitoringService) requireMentor(ctx context.Context, cohortID, userID uint) error {
	role, err := s.repo.GetActiveRole(ctx, cohortID, userID)
	if err != nil {
		return err
	}
	if role != entity.RoleMentor {
		return response.ErrForbidden
	}
	return nil
}

func (s *monitoringService) loadSnapshot(ctx context.Context, cohortID uint, memberSearch string) (*cohortSnapshot, error) {
	materials, err := s.repo.ListMaterials(ctx, cohortID, "")
	if err != nil {
		return nil, err
	}
	members, err := s.repo.ListActiveMembers(ctx, cohortID, memberSearch)
	if err != nil {
		return nil, err
	}
	progresses, err := s.repo.ListCohortProgress(ctx, cohortID)
	if err != nil {
		return nil, err
	}

	idx := make(map[progressKey]*entity.Progress, len(progresses))
	for i := range progresses {
		p := &progresses[i]
		idx[progressKey{p.UserID, p.MaterialID}] = p
	}
	return &cohortSnapshot{materials: materials, members: members, progress: idx}, nil
}

func statusOf(p *entity.Progress) string {
	if p == nil {
		return entity.StatusNotStarted
	}
	return p.Status
}

func isLateProgress(m *entity.Material, p *entity.Progress, now time.Time) bool {
	return statusOf(p) != entity.StatusCompleted && isPastDeadline(m, now)
}

func toMatrixMaterial(m *entity.Material) dto.MatrixMaterial {
	return dto.MatrixMaterial{ID: m.ID, Title: m.Title, Position: m.Position, Deadline: m.Deadline}
}

func reviewerIDsOf(progresses []*entity.Progress) []uint {
	seen := map[uint]struct{}{}
	ids := make([]uint, 0)
	for _, p := range progresses {
		if p == nil || p.ReviewerID == nil {
			continue
		}
		if _, ok := seen[*p.ReviewerID]; ok {
			continue
		}
		seen[*p.ReviewerID] = struct{}{}
		ids = append(ids, *p.ReviewerID)
	}
	return ids
}

// ───────────────── Dashboard ─────────────────

func (s *monitoringService) Dashboard(ctx context.Context, userID, cohortID uint) (*dto.DashboardResponse, error) {
	if err := s.requireMentor(ctx, cohortID, userID); err != nil {
		return nil, err
	}

	snap, err := s.loadSnapshot(ctx, cohortID, "")
	if err != nil {
		return nil, err
	}

	now := s.now()
	res := &dto.DashboardResponse{
		ActiveMemberCount: len(snap.members),
		MaterialCount:     len(snap.materials),
	}

	// hanya pasangan (member aktif × materi yang masih ada) yang dihitung
	for _, mem := range snap.members {
		for i := range snap.materials {
			m := &snap.materials[i]
			p := snap.progress[progressKey{mem.UserID, m.ID}]

			switch statusOf(p) {
			case entity.StatusSubmitted:
				res.PendingReviewCount++
			case entity.StatusCompleted:
				res.CompletedCount++
			}
			if isLateProgress(m, p, now) {
				res.LateCount++
			}
		}
	}

	if total := len(snap.members) * len(snap.materials); total > 0 {
		res.CompletionRate = float64(res.CompletedCount) / float64(total)
	}
	return res, nil
}

// ───────────────── Matrix ─────────────────

func (s *monitoringService) Matrix(ctx context.Context, userID, cohortID uint, q dto.MatrixQuery) (*dto.MatrixResponse, error) {
	if err := s.requireMentor(ctx, cohortID, userID); err != nil {
		return nil, err
	}

	snap, err := s.loadSnapshot(ctx, cohortID, strings.TrimSpace(q.Q))
	if err != nil {
		return nil, err
	}

	// kolom yang ditampilkan: semua materi, atau satu materi kalau material_id diisi
	cols := snap.materials
	if q.MaterialID != 0 {
		cols = nil
		for i := range snap.materials {
			if snap.materials[i].ID == q.MaterialID {
				cols = snap.materials[i : i+1]
				break
			}
		}
		if cols == nil {
			return nil, response.ErrNotFound
		}
	}

	now := s.now()
	res := &dto.MatrixResponse{
		Materials: make([]dto.MatrixMaterial, 0, len(cols)),
		Rows:      make([]dto.MatrixRow, 0, len(snap.members)),
	}
	for i := range cols {
		res.Materials = append(res.Materials, toMatrixMaterial(&cols[i]))
	}

	for _, mem := range snap.members {
		row := dto.MatrixRow{UserID: mem.UserID, Name: mem.Name, Cells: make([]dto.MatrixCell, 0, len(cols))}

		// progres pribadi dihitung dari SEMUA materi, bukan hanya kolom yang difilter
		for i := range snap.materials {
			if statusOf(snap.progress[progressKey{mem.UserID, snap.materials[i].ID}]) == entity.StatusCompleted {
				row.Completed++
			}
		}
		if len(snap.materials) > 0 {
			row.CompletionRate = float64(row.Completed) / float64(len(snap.materials))
		}

		matchStatus := q.Status == ""
		for i := range cols {
			m := &cols[i]
			p := snap.progress[progressKey{mem.UserID, m.ID}]

			cell := dto.MatrixCell{MaterialID: m.ID, Status: statusOf(p), IsLate: isLateProgress(m, p, now)}
			if p != nil {
				id := p.ID
				cell.ProgressID = &id
			}
			if cell.Status == q.Status {
				matchStatus = true
			}
			row.Cells = append(row.Cells, cell)
		}

		if matchStatus {
			res.Rows = append(res.Rows, row)
		}
	}
	return res, nil
}

// ───────────────── hasil belajar per member ─────────────────

func (s *monitoringService) MemberProgress(ctx context.Context, userID, cohortID, memberID uint) (*dto.MemberProgressResponse, error) {
	role, err := s.repo.GetActiveRole(ctx, cohortID, userID)
	if err != nil {
		return nil, err
	}
	// Mentor boleh lihat siapa saja; member hanya boleh lihat miliknya sendiri
	if role != entity.RoleMentor && userID != memberID {
		return nil, response.ErrForbidden
	}

	member, err := s.repo.GetMember(ctx, cohortID, memberID) // termasuk yang sudah keluar
	if err != nil {
		return nil, err
	}

	materials, err := s.repo.ListMaterials(ctx, cohortID, "")
	if err != nil {
		return nil, err
	}
	all, err := s.repo.ListCohortProgress(ctx, cohortID)
	if err != nil {
		return nil, err
	}

	byMaterial := make(map[uint]*entity.Progress)
	mine := make([]*entity.Progress, 0)
	for i := range all {
		if all[i].UserID == memberID {
			byMaterial[all[i].MaterialID] = &all[i]
			mine = append(mine, &all[i])
		}
	}

	names, err := s.repo.GetUserNames(ctx, reviewerIDsOf(mine))
	if err != nil {
		return nil, err
	}

	now := s.now()
	res := &dto.MemberProgressResponse{
		Member: dto.MemberBrief{UserID: member.UserID, Name: member.Name, Role: member.Role, Status: member.Status},
		Items:  make([]dto.MemberProgressItem, 0, len(materials)),
	}
	res.Summary.TotalMaterials = len(materials)

	for i := range materials {
		m := &materials[i]
		p := byMaterial[m.ID]

		switch statusOf(p) {
		case entity.StatusNotStarted:
			res.Summary.NotStarted++
		case entity.StatusInProgress:
			res.Summary.InProgress++
		case entity.StatusSubmitted:
			res.Summary.Submitted++
		case entity.StatusCompleted:
			res.Summary.Completed++
		}
		if isLateProgress(m, p, now) {
			res.Summary.Late++
		}

		res.Items = append(res.Items, dto.MemberProgressItem{
			Material: toMatrixMaterial(m),
			Progress: toProgressResponse(m, p, memberID, now, names),
		})
	}

	if res.Summary.TotalMaterials > 0 {
		res.Summary.CompletionRate = float64(res.Summary.Completed) / float64(res.Summary.TotalMaterials)
	}
	return res, nil
}
