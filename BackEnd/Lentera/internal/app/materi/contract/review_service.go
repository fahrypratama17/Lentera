package contract

import (
	"context"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/dto"
)

type ReviewService interface {
	ListReviews(ctx context.Context, userID, cohortID uint, q dto.ListReviewsQuery) ([]dto.ReviewItemResponse, error)
	GetReviewItem(ctx context.Context, userID, cohortID, progressID uint) (*dto.ReviewItemResponse, error)

	Approve(ctx context.Context, userID, cohortID, progressID uint, req dto.ApproveRequest) (*dto.ReviewItemResponse, error)
	RequestRevision(ctx context.Context, userID, cohortID, progressID uint, req dto.ReviewNoteRequest) (*dto.ReviewItemResponse, error)
	Reopen(ctx context.Context, userID, cohortID, progressID uint, req dto.ReviewNoteRequest) (*dto.ReviewItemResponse, error)
}
