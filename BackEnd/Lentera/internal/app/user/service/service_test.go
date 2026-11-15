package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/user/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/user/entity"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/user/service"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

type mockUserRepo struct {
	users  map[string]*entity.User
	byID   map[uint]*entity.User
	nextID uint
}

func newMockUserRepo() *mockUserRepo {
	return &mockUserRepo{
		users:  make(map[string]*entity.User),
		byID:   make(map[uint]*entity.User),
		nextID: 1,
	}
}

func (m *mockUserRepo) Create(ctx context.Context, u *entity.User) error {
	if _, exists := m.users[u.Email]; exists {
		return errors.New("duplicate email")
	}
	u.ID = m.nextID
	m.nextID++
	m.users[u.Email] = u
	m.byID[u.ID] = u
	return nil
}

func (m *mockUserRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	u, exists := m.users[email]
	if !exists {
		return nil, errors.New("record not found")
	}
	return u, nil
}

func (m *mockUserRepo) FindByID(ctx context.Context, id uint) (*entity.User, error) {
	u, exists := m.byID[id]
	if !exists {
		return nil, errors.New("record not found")
	}
	return u, nil
}

func TestRegister_SuccessAndDuplicateRejected(t *testing.T) {
	repo := newMockUserRepo()
	svc := service.New(repo, "rahasia-testing")
	ctx := context.Background()

	res, err := svc.Register(ctx, dto.RegisterRequest{
		Name:     "Nadia",
		Email:    "Nadia@Mail.com",
		Password: "password123",
	})
	if err != nil {
		t.Fatalf("harusnya berhasil daftar, malah error: %v", err)
	}
	if res.Email != "nadia@mail.com" {
		t.Errorf("email harus dinormalisasi ke huruf kecil, dapat: %s", res.Email)
	}

	_, err = svc.Register(ctx, dto.RegisterRequest{
		Name:     "Nadia Dua",
		Email:    "nadia@mail.com",
		Password: "password123",
	})
	if !errors.Is(err, response.ErrConflict) {
		t.Fatalf("harusnya ditolak dengan ErrConflict, dapat: %v", err)
	}
}

func TestLogin_SuccessAndWrongCredentialsRejected(t *testing.T) {
	repo := newMockUserRepo()
	svc := service.New(repo, "rahasia-testing")
	ctx := context.Background()

	_, _ = svc.Register(ctx, dto.RegisterRequest{
		Name:     "Bagas",
		Email:    "bagas@mail.com",
		Password: "password123",
	})

	authRes, err := svc.Login(ctx, dto.LoginRequest{
		Email:    "bagas@mail.com",
		Password: "password123",
	})
	if err != nil {
		t.Fatalf("login valid harusnya berhasil, dapat error: %v", err)
	}
	if authRes.Token == "" {
		t.Fatal("token JWT tidak boleh kosong")
	}

	_, err = svc.Login(ctx, dto.LoginRequest{
		Email:    "bagas@mail.com",
		Password: "passwordsalah",
	})
	if !errors.Is(err, response.ErrUnauthorized) {
		t.Fatalf("password salah harusnya ditolak ErrUnauthorized, dapat: %v", err)
	}

	_, err = svc.Login(ctx, dto.LoginRequest{
		Email:    "tidakada@mail.com",
		Password: "password123",
	})
	if !errors.Is(err, response.ErrUnauthorized) {
		t.Fatalf("email tidak terdaftar harusnya ditolak ErrUnauthorized, dapat: %v", err)
	}
}

func TestMe_SuccessAndNotFound(t *testing.T) {
	repo := newMockUserRepo()
	svc := service.New(repo, "rahasia-testing")
	ctx := context.Background()

	created, _ := svc.Register(ctx, dto.RegisterRequest{
		Name:     "Sinta",
		Email:    "sinta@mail.com",
		Password: "password123",
	})

	profile, err := svc.Me(ctx, created.ID)
	if err != nil {
		t.Fatalf("harusnya berhasil ambil profil, dapat error: %v", err)
	}
	if profile.Name != "Sinta" {
		t.Errorf("nama tidak sesuai, dapat: %s", profile.Name)
	}

	_, err = svc.Me(ctx, 999)
	if !errors.Is(err, response.ErrNotFound) {
		t.Fatalf("user ID tidak ada harusnya ErrNotFound, dapat: %v", err)
	}
}