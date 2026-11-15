package contract

import (
	"context"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/dto"
)

type MaterialService interface {
	ListMaterials(ctx context.Context, userID, cohortID uint, q dto.ListMaterialsQuery) ([]dto.MaterialResponse, error)
	GetMaterial(ctx context.Context, userID, cohortID, materialID uint) (*dto.MaterialResponse, error)
	MySummary(ctx context.Context, userID, cohortID uint) (*dto.PersonalSummary, error)
}
