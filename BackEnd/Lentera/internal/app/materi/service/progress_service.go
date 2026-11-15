package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/entity"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

const (
	DecisionApproved          = "approved"
	DecisionRevisionRequested = "revision_requested"
)

func (s *materialService) StartMaterial(ctx context.Context, userID, cohortID, materialID uint) (*dto.ProgressResponse, error) {
	role, err := s.repo.GetActiveRole(ctx, cohortID, userID)
	if err != nil {
		return nil, err
	}
	if role != entity.RoleMember {
		return nil, response.ErrForbidden
	}

	m, err := s.repo.GetMaterial(ctx, cohortID, materialID)
	if err != nil {
		return nil, err
	}

	p, err := s.repo.GetUserProgress(ctx, materialID, userID)
	if err != nil {
		return nil, err
	}

	now := s.now()

	if p != nil {
		switch p.Status {
		case entity.StatusInProgress:
			// idempoten: kembalikan progres berjalan tanpa mengubah started_at
			names, err := s.reviewerNames(ctx, []entity.Progress{*p})
			if err != nil {
				return nil, err
			}
			res := toProgressResponse(m, p, userID, now, names)
			return &res, nil
		case entity.StatusSubmitted, entity.StatusCompleted:
			return nil, fmt.Errorf("%w: materi sudah dikumpulkan atau selesai", response.ErrConflict)
		}
	} else {
		p = &entity.Progress{
			CohortID:   cohortID,
			MaterialID: materialID,
			UserID:     userID,
		}
	}

	p.Status = entity.StatusInProgress
	p.StartedAt = &now

	if err := s.repo.SaveProgress(ctx, p); err != nil {
		return nil, err
	}

	res := toProgressResponse(m, p, userID, now, nil)
	s.rt.Broadcast(cohortRoom(cohortID), "progress.updated", res)
	s.notifyDashboard(cohortID)
	return &res, nil
}

func (s *materialService) SubmitMaterial(ctx context.Context, userID, cohortID, materialID uint, req dto.SubmitMaterialRequest) (*dto.ProgressResponse, error) {
	role, err := s.repo.GetActiveRole(ctx, cohortID, userID)
	if err != nil {
		return nil, err
	}
	if role != entity.RoleMember {
		return nil, response.ErrForbidden
	}

	content := strings.TrimSpace(req.Content)
	var workURL *string
	if req.WorkURL != nil && strings.TrimSpace(*req.WorkURL) != "" {
		u := strings.TrimSpace(*req.WorkURL)
		workURL = &u
	}
	if content == "" && workURL == nil {
		return nil, fmt.Errorf("%w: uraian atau tautan hasil wajib diisi", response.ErrBadRequest)
	}

	m, err := s.repo.GetMaterial(ctx, cohortID, materialID)
	if err != nil {
		return nil, err
	}

	p, err := s.repo.GetUserProgress(ctx, materialID, userID)
	if err != nil {
		return nil, err
	}
	if p == nil || p.Status == entity.StatusNotStarted {
		return nil, fmt.Errorf("%w: mulai materi terlebih dahulu sebelum mengumpulkan hasil", response.ErrConflict)
	}
	if p.Status != entity.StatusInProgress {
		return nil, fmt.Errorf("%w: pengumpulan tidak dapat diubah saat status %s", response.ErrConflict, p.Status)
	}

	now := s.now()
	p.Status = entity.StatusSubmitted
	p.SubmissionContent = &content
	p.SubmissionURL = workURL
	p.SubmittedAt = &now

	if err := s.repo.SaveProgress(ctx, p); err != nil {
		return nil, err
	}

	names, err := s.reviewerNames(ctx, []entity.Progress{*p})
	if err != nil {
		return nil, err
	}

	res := toProgressResponse(m, p, userID, now, names)
	s.rt.Broadcast(cohortRoom(cohortID), "progress.updated", res)
	s.notifyDashboard(cohortID)
	return &res, nil
}

func (s *materialService) ReviewSubmission(ctx context.Context, mentorID, cohortID, materialID, targetUserID uint, req dto.ReviewSubmissionRequest) (*dto.ProgressResponse, error) {
	if err := s.requireMentor(ctx, cohortID, mentorID); err != nil {
		return nil, err
	}

	m, err := s.repo.GetMaterial(ctx, cohortID, materialID)
	if err != nil {
		return nil, err
	}

	targetRole, err := s.repo.GetActiveRole(ctx, cohortID, targetUserID)
	if err != nil {
		return nil, err
	}
	if targetRole != entity.RoleMember {
		return nil, fmt.Errorf("%w: hanya progres member yang dapat direview", response.ErrBadRequest)
	}

	p, err := s.repo.GetUserProgress(ctx, materialID, targetUserID)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, fmt.Errorf("%w: member belum memulai materi ini", response.ErrNotFound)
	}

	now := s.now()
	var note *string
	if req.Note != nil && strings.TrimSpace(*req.Note) != "" {
		n := strings.TrimSpace(*req.Note)
		note = &n
	}

	switch req.Decision {
	case DecisionApproved:
		if p.Status != entity.StatusSubmitted {
			return nil, fmt.Errorf("%w: hanya pengumpulan berstatus submitted yang dapat disetujui", response.ErrConflict)
		}
		p.Status = entity.StatusCompleted
		p.CompletedAt = &now
	case DecisionRevisionRequested:
		if p.Status != entity.StatusSubmitted && p.Status != entity.StatusCompleted {
			return nil, fmt.Errorf("%w: pengumpulan belum diajukan untuk direview", response.ErrConflict)
		}
		if note == nil {
			return nil, fmt.Errorf("%w: catatan revisi wajib diisi", response.ErrBadRequest)
		}
		p.Status = entity.StatusInProgress
		p.CompletedAt = nil
	default:
		return nil, fmt.Errorf("%w: keputusan review tidak valid", response.ErrBadRequest)
	}

	decision := req.Decision
	p.ReviewDecision = &decision
	p.ReviewNote = note
	p.ReviewerID = &mentorID
	p.ReviewedAt = &now

	if err := s.repo.SaveProgress(ctx, p); err != nil {
		return nil, err
	}

	names, err := s.reviewerNames(ctx, []entity.Progress{*p})
	if err != nil {
		return nil, err
	}

	res := toProgressResponse(m, p, targetUserID, now, names)
	s.rt.Broadcast(cohortRoom(cohortID), "progress.updated", res)
	s.notifyDashboard(cohortID)
	return &res, nil
}