package service

import (
	"context"
	"strings"
	"time"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/user/contract"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/user/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/user/entity"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/jwt"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/utils"
)

type service struct {
	repo      contract.UserRepository
	jwtSecret string
}

func New(repo contract.UserRepository, jwtSecret string) contract.UserService {
	return &service{repo: repo, jwtSecret: jwtSecret}
}

func (s *service) Register(ctx context.Context, req dto.RegisterRequest) (*dto.UserResponse, error) {
	email := strings.ToLower(strings.TrimSpace(req.Email))

	if existing, _ := s.repo.FindByEmail(ctx, email); existing != nil {
		return nil, response.ErrConflict
	}

	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Name:         strings.TrimSpace(req.Name),
		Email:        email,
		PasswordHash: hash,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, response.ErrConflict
	}

	return &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *service) Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, error) {
	email := strings.ToLower(strings.TrimSpace(req.Email))

	u, err := s.repo.FindByEmail(ctx, email)
	if err != nil || u == nil {
		return nil, response.ErrUnauthorized
	}

	if !utils.CheckPassword(req.Password, u.PasswordHash) {
		return nil, response.ErrUnauthorized
	}

	token, err := jwt.Sign(s.jwtSecret, u.ID, "", 24*time.Hour)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Token: token,
		User: dto.UserResponse{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
		},
	}, nil
}

func (s *service) Me(ctx context.Context, userID uint) (*dto.UserResponse, error) {
	u, err := s.repo.FindByID(ctx, userID)
	if err != nil || u == nil {
		return nil, response.ErrNotFound
	}

	return &dto.UserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}, nil
}