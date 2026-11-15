package repository

import (
	"context"
	"errors"
	"strings"

	"gorm.io/gorm"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/contract"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/entity"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

type materialRepository struct {
	db *gorm.DB
}

func NewMaterialRepository(db *gorm.DB) contract.MaterialRepository {
	return &materialRepository{db: db}
}

// ⚠️ SESUAIKAN: nama tabel & kolom membership harus sama dengan punya BE 1
func (r *materialRepository) GetActiveRole(ctx context.Context, cohortID, userID uint) (string, error) {
	var role string
	err := r.db.WithContext(ctx).
		Table("memberships").
		Select("role").
		Where("cohort_id = ? AND user_id = ? AND status = ?", cohortID, userID, "active").
		Limit(1).
		Scan(&role).Error
	if err != nil {
		return "", err
	}
	if role == "" {
		return "", response.ErrForbidden
	}
	return role, nil
}

func (r *materialRepository) ListMaterials(ctx context.Context, cohortID uint, search string) ([]entity.Material, error) {
	tx := r.db.WithContext(ctx).Where("cohort_id = ?", cohortID)
	if search != "" {
		tx = tx.Where("title ILIKE ?", "%"+escapeLike(search)+"%")
	}

	var out []entity.Material
	err := tx.Order("position ASC, id ASC").Find(&out).Error
	return out, err
}

func (r *materialRepository) GetMaterial(ctx context.Context, cohortID, materialID uint) (*entity.Material, error) {
	var m entity.Material
	err := r.db.WithContext(ctx).
		Where("id = ? AND cohort_id = ?", materialID, cohortID).
		First(&m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, response.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *materialRepository) ListUserProgress(ctx context.Context, cohortID, userID uint) ([]entity.Progress, error) {
	var out []entity.Progress
	err := r.db.WithContext(ctx).
		Where("cohort_id = ? AND user_id = ?", cohortID, userID).
		Find(&out).Error
	return out, err
}

// return (nil, nil) kalau belum ada progres = not_started
func (r *materialRepository) GetUserProgress(ctx context.Context, materialID, userID uint) (*entity.Progress, error) {
	var p entity.Progress
	err := r.db.WithContext(ctx).
		Where("material_id = ? AND user_id = ?", materialID, userID).
		First(&p).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *materialRepository) GetUserNames(ctx context.Context, ids []uint) (map[uint]string, error) {
	out := make(map[uint]string, len(ids))
	if len(ids) == 0 {
		return out, nil
	}

	var rows []struct {
		ID   uint
		Name string
	}
	err := r.db.WithContext(ctx).
		Table("users").
		Select("id, name").
		Where("id IN ?", ids).
		Scan(&rows).Error
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		out[row.ID] = row.Name
	}
	return out, nil
}

// cegah input user % atau _ jadi wildcard LIKE
func escapeLike(s string) string {
	return strings.NewReplacer(`\`, `\\`, `%`, `\%`, `_`, `\_`).Replace(s)
}
