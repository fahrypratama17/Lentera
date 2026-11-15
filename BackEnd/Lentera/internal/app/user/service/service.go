package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/user/contract"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/user/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/user/entity"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/jwt"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/utils"
)

const (
	accessTokenTTL  = 15 * time.Minute
	refreshTokenTTL = 7 * 24 * time.Hour
)

type service struct {
	repo      contract.UserRepository
	jwtSecret string
}

func New(repo contract.UserRepository, jwtSecret string) contract.UserService {
	return &service{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (s *service) Register(ctx context.Context, req dto.RegisterRequest) (*dto.AuthResponse, error) {
	email := strings.ToLower(strings.TrimSpace(req.Email))

	existing, err := s.repo.FindByEmail(ctx, email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existing != nil {
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
		return nil, err
	}

	return s.issueTokens(ctx, *user)
}

func (s *service) Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, error) {
	email := strings.ToLower(strings.TrimSpace(req.Email))

	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrUnauthorized
		}
		return nil, err
	}

	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		return nil, response.ErrUnauthorized
	}

	return s.issueTokens(ctx, *user)
}

func (s *service) RefreshToken(ctx context.Context, req dto.RefreshTokenRequest) (*dto.AuthResponse, error) {
	tokenHash := jwt.HashToken(strings.TrimSpace(req.RefreshToken))

	rt, err := s.repo.FindRefreshTokenByHash(ctx, tokenHash)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrUnauthorized
		}
		return nil, err
	}

	if rt.Revoked || time.Now().After(rt.ExpiresAt) {
		return nil, response.ErrUnauthorized
	}

	// Token rotation: cabut token lama sebelum menerbitkan pasangan token baru
	if err := s.repo.RevokeRefreshToken(ctx, tokenHash); err != nil {
		return nil, err
	}

	user, err := s.repo.FindByID(ctx, rt.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrUnauthorized
		}
		return nil, err
	}

	return s.issueTokens(ctx, *user)
}

func (s *service) Logout(ctx context.Context, req dto.RefreshTokenRequest) error {
	tokenHash := jwt.HashToken(strings.TrimSpace(req.RefreshToken))
	return s.repo.RevokeRefreshToken(ctx, tokenHash)
}

func (s *service) Me(ctx context.Context, userID uint) (*dto.UserResponse, error) {
	user, err := s.repo.FindByID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, err
	}

	res := dto.FromUser(*user)
	return &res, nil
}

func (s *service) issueTokens(ctx context.Context, user entity.User) (*dto.AuthResponse, error) {
	accessToken, err := jwt.Sign(s.jwtSecret, user.ID, "", accessTokenTTL)
	if err != nil {
		return nil, err
	}

	rawRefresh, refreshHash, err := jwt.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	rt := &entity.RefreshToken{
		UserID:    user.ID,
		TokenHash: refreshHash,
		ExpiresAt: time.Now().Add(refreshTokenTTL),
	}

	if err := s.repo.CreateRefreshToken(ctx, rt); err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: rawRefresh,
		User:         dto.FromUser(user),
	}, nil
}