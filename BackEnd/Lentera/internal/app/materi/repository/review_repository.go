package repository

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/contract"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/entity"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

func NewReview(db *gorm.DB) contract.ReviewRepository {
	return &materialRepository{db: db}
}

// ⚠️ SESUAIKAN dengan skema memberships milik Mone
// SEMUA anggota cohort (aktif & sudah keluar).
func (r *materialRepository) ListCohortMembers(ctx context.Context, cohortID uint) ([]entity.MemberInfo, error) {
	var out []entity.MemberInfo
	err := r.db.WithContext(ctx).
		Table("memberships AS ms").
		Select("ms.user_id, u.name, ms.role, ms.status").
		Joins("JOIN users u ON u.id = ms.user_id").
		Where("ms.cohort_id = ?", cohortID).
		Scan(&out).Error
	return out, err
}

// Progres dengan status tertentu, opsional dibatasi 1 materi.
// Urut dari yang paling lama dikumpulkan (antrian review).
func (r *materialRepository) ListProgressByStatus(ctx context.Context, cohortID uint, status string, materialID uint) ([]entity.Progress, error) {
	tx := r.db.WithContext(ctx).
		Where("cohort_id = ? AND status = ?", cohortID, status)
	if materialID != 0 {
		tx = tx.Where("material_id = ?", materialID)
	}

	var out []entity.Progress
	err := tx.Order("submitted_at ASC NULLS LAST, id ASC").Find(&out).Error
	return out, err
}

func (r *materialRepository) GetProgress(ctx context.Context, cohortID, progressID uint) (*entity.Progress, error) {
	var p entity.Progress
	err := r.db.WithContext(ctx).
		Where("id = ? AND cohort_id = ?", progressID, cohortID).
		First(&p).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, response.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// Simpan hasil review HANYA kalau status di DB masih sama dengan saat dibaca.
// Mencegah dua mentor me-review progres yang sama bersamaan.
func (r *materialRepository) SaveReview(ctx context.Context, p *entity.Progress, fromStatus string) error {
	res := r.db.WithContext(ctx).
		Model(&entity.Progress{}).
		Where("id = ? AND cohort_id = ? AND status = ?", p.ID, p.CohortID, fromStatus).
		Updates(map[string]any{
			"status":          p.Status,
			"completed_at":    p.CompletedAt,
			"review_decision": p.ReviewDecision,
			"review_note":     p.ReviewNote,
			"reviewer_id":     p.ReviewerID,
			"reviewed_at":     p.ReviewedAt,
		})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("%w: progres sudah direview atau statusnya berubah, muat ulang data", response.ErrConflict)
	}
	return nil
}
