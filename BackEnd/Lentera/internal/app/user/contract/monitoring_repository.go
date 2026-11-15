package contract

import (
	"context"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/entity"
)

type MonitoringRepository interface {
	// dipakai bareng dengan MaterialRepository (implementasinya sama)
	GetActiveRole(ctx context.Context, cohortID, userID uint) (string, error)
	ListMaterials(ctx context.Context, cohortID uint, search string) ([]entity.Material, error)
	GetUserNames(ctx context.Context, ids []uint) (map[uint]string, error)

	// baru: khusus pemantauan
	ListActiveMembers(ctx context.Context, cohortID uint, search string) ([]entity.MemberInfo, error)
	GetMember(ctx context.Context, cohortID, userID uint) (*entity.MemberInfo, error)
	ListCohortProgress(ctx context.Context, cohortID uint) ([]entity.Progress, error)
}
