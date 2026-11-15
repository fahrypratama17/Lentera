package repository

import (
	"context"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/user/contract"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/user/entity"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) contract.UserRepository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, u *entity.User) error {
	// TODO: return r.db.WithContext(ctx).Create(u).Error  (duplikat email -> gorm.ErrDuplicatedKey)
	return response.ErrNotImplemented
}

func (r *repository) FindByID(ctx context.Context, id uint) (*entity.User, error) {
	// TODO
	return nil, response.ErrNotImplemented
}

func (r *repository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	// TODO: Where("email = ?", email).First(&u)
	return nil, response.ErrNotImplemented
}
