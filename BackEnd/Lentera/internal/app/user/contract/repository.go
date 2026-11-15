package contract

import (
	"context"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/user/entity"
)

type UserRepository interface {
	Create(ctx context.Context, u *entity.User) error
	FindByID(ctx context.Context, id uint) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	CreateRefreshToken(ctx context.Context, rt *entity.RefreshToken) error
	FindRefreshTokenByHash(ctx context.Context, tokenHash string) (*entity.RefreshToken, error)
	RevokeRefreshToken(ctx context.Context, tokenHash string) error
}
