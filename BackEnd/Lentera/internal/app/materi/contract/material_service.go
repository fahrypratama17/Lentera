package contract

import (
	"context"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/dto"
)

type MaterialService interface {
	ListMaterials(ctx context.Context, userID, cohortID uint, q dto.ListMaterialsQuery) ([]dto.MaterialResponse, error)
	GetMaterial(ctx context.Context, userID, cohortID, materialID uint) (*dto.MaterialResponse, error)
	MySummary(ctx context.Context, userID, cohortID uint) (*dto.PersonalSummary, error)

	// ── baru: CRUD materi (khusus Mentor) ──
	CreateMaterial(ctx context.Context, userID, cohortID uint, req dto.CreateMaterialRequest) (*dto.MaterialResponse, error)
	UpdateMaterial(ctx context.Context, userID, cohortID, materialID uint, req dto.UpdateMaterialRequest) (*dto.MaterialResponse, error)
	DeleteMaterial(ctx context.Context, userID, cohortID, materialID uint) error
	ReorderMaterials(ctx context.Context, userID, cohortID uint, req dto.ReorderMaterialsRequest) ([]dto.MaterialResponse, error)
}
