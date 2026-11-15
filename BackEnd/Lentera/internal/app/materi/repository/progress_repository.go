package repository

import (
	"context"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/entity"
)

func (r *materialRepository) SaveProgress(ctx context.Context, p *entity.Progress) error {
	return r.db.WithContext(ctx).Save(p).Error
}