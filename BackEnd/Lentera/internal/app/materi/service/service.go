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

type materialService struct {
	repo contract.MaterialRepository
	now  func() time.Time
}

func NewMaterialService(repo contract.MaterialRepository) contract.MaterialService {
	return &materialService{repo: repo, now: time.Now}
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
