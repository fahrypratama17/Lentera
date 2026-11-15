package contract

import (
	"context"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/entity"
)

type MaterialRepository interface {
	GetActiveRole(ctx context.Context, cohortID, userID uint) (string, error)
	ListMaterials(ctx context.Context, cohortID uint, search string) ([]entity.Material, error)
	GetMaterial(ctx context.Context, cohortID, materialID uint) (*entity.Material, error)
	ListUserProgress(ctx context.Context, cohortID, userID uint) ([]entity.Progress, error)
	GetUserProgress(ctx context.Context, materialID, userID uint) (*entity.Progress, error)
	GetUserNames(ctx context.Context, ids []uint) (map[uint]string, error)

	// ── baru: CRUD materi ──
	NextPosition(ctx context.Context, cohortID uint) (int, error)
	CreateMaterial(ctx context.Context, m *entity.Material) error
	UpdateMaterial(ctx context.Context, cohortID, materialID uint, fields map[string]any) error
	DeleteMaterial(ctx context.Context, cohortID, materialID uint) error
	ReorderMaterials(ctx context.Context, cohortID uint, orderedIDs []uint) error
}
