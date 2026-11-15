package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/contract"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/entity"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

type materialRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) contract.MaterialRepository {
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

// ───────────────── CRUD materi ─────────────────

// materi baru selalu di urutan terakhir
func (r *materialRepository) NextPosition(ctx context.Context, cohortID uint) (int, error) {
	var maxPos int
	err := r.db.WithContext(ctx).
		Model(&entity.Material{}).
		Where("cohort_id = ?", cohortID).
		Select("COALESCE(MAX(position), 0)").
		Scan(&maxPos).Error
	if err != nil {
		return 0, err
	}
	return maxPos + 1, nil
}

func (r *materialRepository) CreateMaterial(ctx context.Context, m *entity.Material) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *materialRepository) UpdateMaterial(ctx context.Context, cohortID, materialID uint, fields map[string]any) error {
	if len(fields) == 0 {
		return nil
	}

	// update via map tidak lewat serializer:json, jadi encode manual ke jsonb
	if refs, ok := fields["references"].([]string); ok {
		b, err := json.Marshal(refs)
		if err != nil {
			return err
		}
		fields["references"] = gorm.Expr("?::jsonb", string(b))
	}

	res := r.db.WithContext(ctx).
		Model(&entity.Material{}).
		Where("id = ? AND cohort_id = ?", materialID, cohortID).
		Updates(fields)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return response.ErrNotFound
	}
	return nil
}

func (r *materialRepository) DeleteMaterial(ctx context.Context, cohortID, materialID uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		res := tx.Where("id = ? AND cohort_id = ?", materialID, cohortID).
			Delete(&entity.Material{})
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return response.ErrNotFound
		}

		// jaga-jaga kalau FK cascade belum terpasang
		return tx.Where("material_id = ?", materialID).
			Delete(&entity.Progress{}).Error
	})
}

func (r *materialRepository) ReorderMaterials(ctx context.Context, cohortID uint, orderedIDs []uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// kunci baris materi cohort ini supaya dua reorder bersamaan tidak saling timpa
		var existing []uint
		err := tx.Model(&entity.Material{}).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("cohort_id = ?", cohortID).
			Pluck("id", &existing).Error
		if err != nil {
			return err
		}

		if err := sameIDSet(existing, orderedIDs); err != nil {
			return err
		}

		for i, id := range orderedIDs {
			err := tx.Model(&entity.Material{}).
				Where("id = ? AND cohort_id = ?", id, cohortID).
				Update("position", i+1).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// orderedIDs harus berisi SEMUA materi cohort, tanpa kurang/lebih/duplikat
func sameIDSet(existing, ordered []uint) error {
	if len(existing) != len(ordered) {
		return fmt.Errorf("%w: material_ids harus berisi semua %d materi cohort", response.ErrBadRequest, len(existing))
	}

	valid := make(map[uint]bool, len(existing))
	for _, id := range existing {
		valid[id] = true
	}

	seen := make(map[uint]bool, len(ordered))
	for _, id := range ordered {
		if !valid[id] {
			return fmt.Errorf("%w: materi %d bukan milik cohort ini", response.ErrBadRequest, id)
		}
		if seen[id] {
			return fmt.Errorf("%w: materi %d duplikat", response.ErrBadRequest, id)
		}
		seen[id] = true
	}
	return nil
}
