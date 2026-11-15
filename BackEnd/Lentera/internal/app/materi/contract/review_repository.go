package contract

import (
	"context"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/entity"
)

type ReviewRepository interface {
	// dipakai bareng (implementasi sudah ada)
	GetActiveRole(ctx context.Context, cohortID, userID uint) (string, error)
	ListMaterials(ctx context.Context, cohortID uint, search string) ([]entity.Material, error)
	GetUserNames(ctx context.Context, ids []uint) (map[uint]string, error)

	// baru: khusus review
	ListCohortMembers(ctx context.Context, cohortID uint) ([]entity.MemberInfo, error)
	ListProgressByStatus(ctx context.Context, cohortID uint, status string, materialID uint) ([]entity.Progress, error)
	GetProgress(ctx context.Context, cohortID, progressID uint) (*entity.Progress, error)
	SaveReview(ctx context.Context, p *entity.Progress, fromStatus string) error
}
