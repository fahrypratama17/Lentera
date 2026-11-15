package service

import (
	"time"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/entity"
)

func isPastDeadline(m *entity.Material, now time.Time) bool {
	return m.Deadline != nil && now.After(*m.Deadline)
}

func toMaterialResponse(m *entity.Material, p *dto.ProgressResponse, now time.Time) dto.MaterialResponse {
	refs := m.References
	if refs == nil {
		refs = []string{}
	}
	return dto.MaterialResponse{
		ID:             m.ID,
		CohortID:       m.CohortID,
		Title:          m.Title,
		Guide:          m.Guide,
		References:     refs,
		Deadline:       m.Deadline,
		IsPastDeadline: isPastDeadline(m, now),
		Position:       m.Position,
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
		MyProgress:     p,
	}
}

func toProgressResponse(m *entity.Material, p *entity.Progress, userID uint, now time.Time, names map[uint]string) dto.ProgressResponse {
	// belum ada row = belum pernah dimulai
	if p == nil {
		return dto.ProgressResponse{
			UserID:     userID,
			MaterialID: m.ID,
			Status:     entity.StatusNotStarted,
			IsLate:     isPastDeadline(m, now),
		}
	}

	res := dto.ProgressResponse{
		ID:          &p.ID,
		UserID:      p.UserID,
		MaterialID:  p.MaterialID,
		Status:      p.Status,
		IsLate:      p.Status != entity.StatusCompleted && isPastDeadline(m, now),
		StartedAt:   p.StartedAt,
		CompletedAt: p.CompletedAt,
		UpdatedAt:   &p.UpdatedAt,
	}

	if p.SubmittedAt != nil && p.SubmissionContent != nil {
		res.Submission = &dto.SubmissionResponse{
			Content:     *p.SubmissionContent,
			WorkURL:     p.SubmissionURL,
			SubmittedAt: *p.SubmittedAt,
			IsLate:      m.Deadline != nil && p.SubmittedAt.After(*m.Deadline),
		}
	}

	if p.ReviewDecision != nil && p.ReviewerID != nil && p.ReviewedAt != nil {
		res.LastReview = &dto.ReviewResponse{
			Decision:   *p.ReviewDecision,
			Note:       p.ReviewNote,
			Reviewer:   dto.UserBrief{ID: *p.ReviewerID, Name: names[*p.ReviewerID]},
			ReviewedAt: *p.ReviewedAt,
		}
	}
	return res
}
