package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/contract"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/entity"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

// Struct yang sama dengan repository materi, cukup tambah method.
// GetActiveRole, ListMaterials, GetUserNames sudah ada di material_repository.go.
func NewMonitoring(db *gorm.DB) contract.MonitoringRepository {
	return &materialRepository{db: db}
}

// ⚠️ SESUAIKAN dengan skema memberships milik Mone
// Member aktif (role member) di cohort, opsional cari berdasarkan nama.
func (r *materialRepository) ListActiveMembers(ctx context.Context, cohortID uint, search string) ([]entity.MemberInfo, error) {
	tx := r.db.WithContext(ctx).
		Table("memberships AS ms").
		Select("ms.user_id, u.name, ms.role, ms.status").
		Joins("JOIN users u ON u.id = ms.user_id").
		Where("ms.cohort_id = ? AND ms.status = ? AND ms.role = ?", cohortID, "active", entity.RoleMember)

	if search != "" {
		tx = tx.Where("u.name ILIKE ?", "%"+escapeLike(search)+"%")
	}

	var out []entity.MemberInfo
	err := tx.Order("u.name ASC, ms.user_id ASC").Scan(&out).Error
	return out, err
}

// ⚠️ SESUAIKAN dengan skema memberships milik Mone
// Satu anggota cohort, TERMASUK yang sudah keluar (status = left).
func (r *materialRepository) GetMember(ctx context.Context, cohortID, userID uint) (*entity.MemberInfo, error) {
	var out []entity.MemberInfo
	err := r.db.WithContext(ctx).
		Table("memberships AS ms").
		Select("ms.user_id, u.name, ms.role, ms.status").
		Joins("JOIN users u ON u.id = ms.user_id").
		Where("ms.cohort_id = ? AND ms.user_id = ?", cohortID, userID).
		Limit(1).
		Scan(&out).Error
	if err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, response.ErrNotFound
	}
	return &out[0], nil
}

// Semua progres di cohort dalam SATU query (hindari N+1 query per member).
func (r *materialRepository) ListCohortProgress(ctx context.Context, cohortID uint) ([]entity.Progress, error) {
	var out []entity.Progress
	err := r.db.WithContext(ctx).
		Where("cohort_id = ?", cohortID).
		Find(&out).Error
	return out, err
}
