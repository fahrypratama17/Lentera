package repository

import (
	"context"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/user/contract"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/user/entity"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) contract.UserRepository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, u *entity.User) error {
	return r.db.WithContext(ctx).Create(u).Error
}

func (r *repository) FindByID(ctx context.Context, id uint) (*entity.User, error) {
	var u entity.User
	if err := r.db.WithContext(ctx).First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *repository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var u entity.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *repository) CreateRefreshToken(ctx context.Context, rt *entity.RefreshToken) error {
	return r.db.WithContext(ctx).Create(rt).Error
}

func (r *repository) FindRefreshTokenByHash(ctx context.Context, tokenHash string) (*entity.RefreshToken, error) {
	var rt entity.RefreshToken
	err := r.db.WithContext(ctx).Where("token_hash = ?", tokenHash).First(&rt).Error
	if err != nil {
		return nil, err
	}
	return &rt, nil
}

func (r *repository) RevokeRefreshToken(ctx context.Context, tokenHash string) error {
	return r.db.WithContext(ctx).
		Model(&entity.RefreshToken{}).
		Where("token_hash = ?", tokenHash).
		Update("revoked", true).Error
}