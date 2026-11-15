package service_test

import (
	"context"
	"testing"
	"time"

	"gorm.io/gorm"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/user/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/user/entity"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/user/service"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/jwt"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

type mockUserRepo struct {
	users         map[string]*entity.User
	byID          map[uint]*entity.User
	refreshTokens map[string]*entity.RefreshToken
	nextID        uint
	nextRTID      uint
}

func newMockUserRepo() *mockUserRepo {
	return &mockUserRepo{
		users:         make(map[string]*entity.User),
		byID:          make(map[uint]*entity.User),
		refreshTokens: make(map[string]*entity.RefreshToken),
		nextID:        1,
		nextRTID:      1,
	}
}

func (m *mockUserRepo) Create(ctx context.Context, u *entity.User) error {
	if _, exists := m.users[u.Email]; exists {
		return gorm.ErrDuplicatedKey
	}
	u.ID = m.nextID
	m.nextID++
	m.users[u.Email] = u
	m.byID[u.ID] = u
	return nil
}

func (m *mockUserRepo) FindByID(ctx context.Context, id uint) (*entity.User, error) {
	u, exists := m.byID[id]
	if !exists {
		return nil, gorm.ErrRecordNotFound
	}
	return u, nil
}

func (m *mockUserRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	u, exists := m.users[email]
	if !exists {
		return nil, gorm.ErrRecordNotFound
	}
	return u, nil
}

func (m *mockUserRepo) CreateRefreshToken(ctx context.Context, rt *entity.RefreshToken) error {
	rt.ID = m.nextRTID
	m.nextRTID++
	m.refreshTokens[rt.TokenHash] = rt
	return nil
}

func (m *mockUserRepo) FindRefreshTokenByHash(ctx context.Context, tokenHash string) (*entity.RefreshToken, error) {
	rt, exists := m.refreshTokens[tokenHash]
	if !exists {
		return nil, gorm.ErrRecordNotFound
	}
	return rt, nil
}

func (m *mockUserRepo) RevokeRefreshToken(ctx context.Context, tokenHash string) error {
	rt, exists := m.refreshTokens[tokenHash]
	if !exists {
		return gorm.ErrRecordNotFound
	}
	rt.Revoked = true
	return nil
}

func TestRegister_SuccessAndDuplicateRejected(t *testing.T) {
	ctx := context.Background()
	svc := service.New(newMockUserRepo(), "test-secret")

	res, err := svc.Register(ctx, dto.RegisterRequest{
		Name:     "Mone",
		Email:    "mone@example.com",
		Password: "password123",
	})
	if err != nil {
		t.Fatalf("register valid harusnya berhasil, dapat error: %v", err)
	}
	if res.AccessToken == "" || res.RefreshToken == "" {
		t.Fatal("access_token dan refresh_token tidak boleh kosong")
	}

	_, err = svc.Register(ctx, dto.RegisterRequest{
		Name:     "Mone Dua",
		Email:    "MONE@example.com",
		Password: "password123",
	})
	if err != response.ErrConflict {
		t.Fatalf("duplikat email harusnya ErrConflict, dapat: %v", err)
	}
}

func TestLogin_SuccessAndWrongCredentialsRejected(t *testing.T) {
	ctx := context.Background()
	svc := service.New(newMockUserRepo(), "test-secret")

	_, _ = svc.Register(ctx, dto.RegisterRequest{
		Name:     "Mone",
		Email:    "mone@example.com",
		Password: "password123",
	})

	res, err := svc.Login(ctx, dto.LoginRequest{
		Email:    "mone@example.com",
		Password: "password123",
	})
	if err != nil {
		t.Fatalf("login valid harusnya berhasil, dapat error: %v", err)
	}
	if res.AccessToken == "" || res.RefreshToken == "" {
		t.Fatal("access_token dan refresh_token tidak boleh kosong saat login")
	}

	_, err = svc.Login(ctx, dto.LoginRequest{
		Email:    "mone@example.com",
		Password: "passwordsalah",
	})
	if err != response.ErrUnauthorized {
		t.Fatalf("password salah harusnya ErrUnauthorized, dapat: %v", err)
	}
}

func TestRefreshToken_RotationExpiredAndLogout(t *testing.T) {
	ctx := context.Background()
	repo := newMockUserRepo()
	svc := service.New(repo, "test-secret")

	regRes, err := svc.Register(ctx, dto.RegisterRequest{
		Name:     "Mone",
		Email:    "mone@example.com",
		Password: "password123",
	})
	if err != nil {
		t.Fatalf("gagal register: %v", err)
	}

	refRes, err := svc.RefreshToken(ctx, dto.RefreshTokenRequest{
		RefreshToken: regRes.RefreshToken,
	})
	if err != nil {
		t.Fatalf("refresh token valid harusnya berhasil, dapat error: %v", err)
	}
	if refRes.AccessToken == "" || refRes.RefreshToken == "" {
		t.Fatal("token baru hasil refresh tidak boleh kosong")
	}
	if refRes.RefreshToken == regRes.RefreshToken {
		t.Fatal("refresh token baru harus berbeda dari token lama (token rotation)")
	}

	_, err = svc.RefreshToken(ctx, dto.RefreshTokenRequest{
		RefreshToken: regRes.RefreshToken,
	})
	if err != response.ErrUnauthorized {
		t.Fatalf("refresh token lama yang sudah dipakai harusnya ErrUnauthorized, dapat: %v", err)
	}

	if err := svc.Logout(ctx, dto.RefreshTokenRequest{RefreshToken: refRes.RefreshToken}); err != nil {
		t.Fatalf("logout valid harusnya berhasil, dapat error: %v", err)
	}
	_, err = svc.RefreshToken(ctx, dto.RefreshTokenRequest{
		RefreshToken: refRes.RefreshToken,
	})
	if err != response.ErrUnauthorized {
		t.Fatalf("refresh token setelah logout harusnya ErrUnauthorized, dapat: %v", err)
	}

	loginRes, _ := svc.Login(ctx, dto.LoginRequest{
		Email:    "mone@example.com",
		Password: "password123",
	})
	expiredHash := jwt.HashToken(loginRes.RefreshToken)
	repo.refreshTokens[expiredHash].ExpiresAt = time.Now().Add(-1 * time.Minute)

	_, err = svc.RefreshToken(ctx, dto.RefreshTokenRequest{
		RefreshToken: loginRes.RefreshToken,
	})
	if err != response.ErrUnauthorized {
		t.Fatalf("refresh token expired harusnya ErrUnauthorized, dapat: %v", err)
	}
}

func TestMe_SuccessAndNotFound(t *testing.T) {
	ctx := context.Background()
	svc := service.New(newMockUserRepo(), "test-secret")

	regRes, _ := svc.Register(ctx, dto.RegisterRequest{
		Name:     "Mone",
		Email:    "mone@example.com",
		Password: "password123",
	})

	me, err := svc.Me(ctx, regRes.User.ID)
	if err != nil {
		t.Fatalf("Me valid harusnya berhasil, dapat error: %v", err)
	}
	if me.Email != "mone@example.com" {
		t.Fatalf("email tidak sesuai: got %s", me.Email)
	}

	_, err = svc.Me(ctx, 9999)
	if err != response.ErrNotFound {
		t.Fatalf("user ID tidak ada harusnya ErrNotFound, dapat: %v", err)
	}
}