package service_test

import (
	"context"
	"testing"
	"time"

	"gorm.io/gorm"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/cohort/contract"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/cohort/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/cohort/entity"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/cohort/service"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

type mockCohortRepo struct {
	users       map[uint]*contract.UserBasic
	usersEmail  map[string]*contract.UserBasic
	cohorts     map[uint]*entity.Cohort
	memberships map[uint]map[uint]*entity.Membership
	invitations map[uint]*entity.Invitation
	nextID      uint
}

func newMockCohortRepo() *mockCohortRepo {
	return &mockCohortRepo{
		users:       make(map[uint]*contract.UserBasic),
		usersEmail:  make(map[string]*contract.UserBasic),
		cohorts:     make(map[uint]*entity.Cohort),
		memberships: make(map[uint]map[uint]*entity.Membership),
		invitations: make(map[uint]*entity.Invitation),
		nextID:      1,
	}
}

func (m *mockCohortRepo) seedUser(id uint, name, email string) {
	u := &contract.UserBasic{ID: id, Name: name, Email: email}
	m.users[id] = u
	m.usersEmail[email] = u
}

func (m *mockCohortRepo) CreateWithOwner(ctx context.Context, cohort *entity.Cohort, owner *entity.Membership) error {
	cohort.ID = m.nextID
	m.nextID++
	m.cohorts[cohort.ID] = cohort

	owner.ID = m.nextID
	m.nextID++
	owner.CohortID = cohort.ID
	if m.memberships[cohort.ID] == nil {
		m.memberships[cohort.ID] = make(map[uint]*entity.Membership)
	}
	m.memberships[cohort.ID][owner.UserID] = owner
	return nil
}

func (m *mockCohortRepo) FindUserCohorts(ctx context.Context, userID uint) ([]dto.CohortResponse, error) {
	var res []dto.CohortResponse
	for cID, uMap := range m.memberships {
		if mem, ok := uMap[userID]; ok && mem.Status == entity.MembershipActive {
			c := m.cohorts[cID]
			res = append(res, dto.CohortResponse{
				ID:          c.ID,
				Name:        c.Name,
				Description: c.Description,
				StartDate:   c.StartDate.Format("2006-01-02"),
				EndDate:     c.EndDate.Format("2006-01-02"),
				Role:        mem.Role,
				JoinedAt:    mem.JoinedAt,
			})
		}
	}
	return res, nil
}

func (m *mockCohortRepo) FindCohortByID(ctx context.Context, cohortID uint) (*entity.Cohort, error) {
	c, ok := m.cohorts[cohortID]
	if !ok {
		return nil, gorm.ErrRecordNotFound
	}
	return c, nil
}

func (m *mockCohortRepo) FindMembership(ctx context.Context, cohortID, userID uint) (*entity.Membership, error) {
	if uMap, ok := m.memberships[cohortID]; ok {
		if mem, exists := uMap[userID]; exists {
			return mem, nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *mockCohortRepo) FindUserByID(ctx context.Context, userID uint) (*contract.UserBasic, error) {
	u, ok := m.users[userID]
	if !ok {
		return nil, gorm.ErrRecordNotFound
	}
	return u, nil
}

func (m *mockCohortRepo) FindUserByEmail(ctx context.Context, email string) (*contract.UserBasic, error) {
	u, ok := m.usersEmail[email]
	if !ok {
		return nil, gorm.ErrRecordNotFound
	}
	return u, nil
}

func (m *mockCohortRepo) CreateInvitation(ctx context.Context, inv *entity.Invitation) error {
	inv.ID = m.nextID
	m.nextID++
	m.invitations[inv.ID] = inv
	return nil
}

func (m *mockCohortRepo) FindPendingInvitation(ctx context.Context, cohortID uint, email string) (*entity.Invitation, error) {
	for _, inv := range m.invitations {
		if inv.CohortID == cohortID && inv.Email == email && inv.Status == entity.InvitationPending {
			return inv, nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *mockCohortRepo) FindInvitationByID(ctx context.Context, invitationID uint) (*entity.Invitation, error) {
	inv, ok := m.invitations[invitationID]
	if !ok {
		return nil, gorm.ErrRecordNotFound
	}
	return inv, nil
}

func (m *mockCohortRepo) FindUserPendingInvitations(ctx context.Context, email string) ([]dto.InvitationResponse, error) {
	var res []dto.InvitationResponse
	for _, inv := range m.invitations {
		if inv.Email == email && inv.Status == entity.InvitationPending {
			res = append(res, dto.InvitationResponse{
				ID:          inv.ID,
				CohortID:    inv.CohortID,
				CohortName:  m.cohorts[inv.CohortID].Name,
				InviterID:   inv.InviterID,
				InviterName: m.users[inv.InviterID].Name,
				Email:       inv.Email,
				Role:        inv.Role,
				Status:      inv.Status,
			})
		}
	}
	return res, nil
}

func (m *mockCohortRepo) FindCohortInvitations(ctx context.Context, cohortID uint) ([]dto.InvitationResponse, error) {
	var res []dto.InvitationResponse
	for _, inv := range m.invitations {
		if inv.CohortID == cohortID {
			res = append(res, dto.InvitationResponse{
				ID:     inv.ID,
				Email:  inv.Email,
				Role:   inv.Role,
				Status: inv.Status,
			})
		}
	}
	return res, nil
}

func (m *mockCohortRepo) UpdateInvitationStatus(ctx context.Context, invitationID uint, status entity.InvitationStatus, respondedAt *time.Time) error {
	inv, ok := m.invitations[invitationID]
	if !ok {
		return gorm.ErrRecordNotFound
	}
	inv.Status = status
	inv.RespondedAt = respondedAt
	return nil
}

func (m *mockCohortRepo) AcceptInvitationTx(ctx context.Context, inv *entity.Invitation, userID uint, now time.Time) error {
	inv.Status = entity.InvitationAccepted
	inv.RespondedAt = &now
	if m.memberships[inv.CohortID] == nil {
		m.memberships[inv.CohortID] = make(map[uint]*entity.Membership)
	}
	m.memberships[inv.CohortID][userID] = &entity.Membership{
		CohortID: inv.CohortID,
		UserID:   userID,
		Role:     inv.Role,
		Status:   entity.MembershipActive,
		JoinedAt: now,
	}
	return nil
}

func TestEmailInvitationFlow_EndToEnd(t *testing.T) {
	ctx := context.Background()
	repo := newMockCohortRepo()
	repo.seedUser(1, "Nadia", "nadia@gmail.com")
	repo.seedUser(2, "Bagas", "bagas@gmail.com")
	repo.seedUser(3, "Rina", "rina@gmail.com")

	svc := service.New(repo)

	cohort, err := svc.CreateCohort(ctx, 1, dto.CreateCohortRequest{
		Name:        "Belajar Web",
		Description: "Cohort Frontend & Backend",
		StartDate:   "2026-11-15",
		EndDate:     "2026-12-15",
	})
	if err != nil {
		t.Fatalf("gagal membuat cohort: %v", err)
	}
	if cohort.Role != entity.RoleMentor {
		t.Fatalf("pembuat cohort harusnya mentor, dapat: %s", cohort.Role)
	}

	_, err = svc.InviteMember(ctx, 1, cohort.ID, dto.CreateInvitationRequest{
		Email: "belumdaftar@gmail.com",
		Role:  entity.RoleMember,
	})
	if err != response.ErrNotFound {
		t.Fatalf("email belum terdaftar harusnya ErrNotFound, dapat: %v", err)
	}

	inv, err := svc.InviteMember(ctx, 1, cohort.ID, dto.CreateInvitationRequest{
		Email: "bagas@gmail.com",
		Role:  entity.RoleMember,
	})
	if err != nil {
		t.Fatalf("gagal mengundang Bagas: %v", err)
	}

	bagasInvs, err := svc.ListMyInvitations(ctx, 2)
	if err != nil || len(bagasInvs) != 1 || bagasInvs[0].ID != inv.ID {
		t.Fatalf("Bagas harusnya melihat 1 undangan pending, dapat: %+v, err: %v", bagasInvs, err)
	}

	if err := svc.RespondInvitation(ctx, 3, inv.ID, dto.RespondInvitationRequest{Action: "accept"}); err != response.ErrForbidden {
		t.Fatalf("Rina tidak boleh menerima undangan Bagas, harusnya ErrForbidden, dapat: %v", err)
	}

	if err := svc.RespondInvitation(ctx, 2, inv.ID, dto.RespondInvitationRequest{Action: "accept"}); err != nil {
		t.Fatalf("Bagas gagal menerima undangan: %v", err)
	}

	bagasCohorts, _ := svc.ListMyCohorts(ctx, 2)
	if len(bagasCohorts) != 1 || bagasCohorts[0].Role != entity.RoleMember {
		t.Fatalf("Bagas harusnya terdaftar aktif di cohort sebagai member, dapat: %+v", bagasCohorts)
	}
}

// Tambahkan method berikut ke mockCohortRepo:
func (m *mockCohortRepo) UpdateCohort(ctx context.Context, cohort *entity.Cohort) error {
	m.cohorts[cohort.ID] = cohort
	return nil
}

func (m *mockCohortRepo) DeleteCohort(ctx context.Context, cohortID uint) error {
	delete(m.cohorts, cohortID)
	delete(m.memberships, cohortID)
	return nil
}

func (m *mockCohortRepo) ListActiveMembers(ctx context.Context, cohortID uint) ([]dto.MemberResponse, error) {
	var res []dto.MemberResponse
	for uID, mem := range m.memberships[cohortID] {
		if mem.Status == entity.MembershipActive {
			u := m.users[uID]
			res = append(res, dto.MemberResponse{
				UserID:   u.ID,
				Name:     u.Name,
				Email:    u.Email,
				Role:     mem.Role,
				JoinedAt: mem.JoinedAt,
			})
		}
	}
	return res, nil
}

func (m *mockCohortRepo) CountActiveMentors(ctx context.Context, cohortID uint) (int64, error) {
	var count int64
	for _, mem := range m.memberships[cohortID] {
		if mem.Status == entity.MembershipActive && mem.Role == entity.RoleMentor {
			count++
		}
	}
	return count, nil
}

func (m *mockCohortRepo) UpdateMemberRole(ctx context.Context, cohortID, userID uint, role entity.CohortRole) error {
	m.memberships[cohortID][userID].Role = role
	return nil
}

func (m *mockCohortRepo) LeaveMembership(ctx context.Context, cohortID, userID uint, leftAt time.Time) error {
	mem := m.memberships[cohortID][userID]
	mem.Status = entity.MembershipLeft
	mem.LeftAt = &leftAt
	return nil
}

func TestMembersAndSettings_LastMentorGuardAndLifecycle(t *testing.T) {
	ctx := context.Background()
	repo := newMockCohortRepo()
	repo.seedUser(1, "Nadia", "nadia@gmail.com")
	repo.seedUser(2, "Bagas", "bagas@gmail.com")

	svc := service.New(repo)
	cohort, _ := svc.CreateCohort(ctx, 1, dto.CreateCohortRequest{
		Name:      "Cohort Backend",
		StartDate: "2026-11-15",
		EndDate:   "2026-12-15",
	})

	inv, _ := svc.InviteMember(ctx, 1, cohort.ID, dto.CreateInvitationRequest{
		Email: "bagas@gmail.com",
		Role:  entity.RoleMember,
	})
	_ = svc.RespondInvitation(ctx, 2, inv.ID, dto.RespondInvitationRequest{Action: "accept"})

	_, err := svc.UpdateCohort(ctx, 2, cohort.ID, dto.UpdateCohortRequest{
		Name:      "CI/CD Backend",
		StartDate: "2026-11-15",
		EndDate:   "2026-12-15",
	})
	if err != response.ErrForbidden {
		t.Fatalf("member tidak boleh edit cohort, harusnya ErrForbidden, dapat: %v", err)
	}

	if err := svc.LeaveCohort(ctx, 1, cohort.ID); err != response.ErrBadRequest {
		t.Fatalf("mentor terakhir harusnya ditolak keluar (ErrBadRequest), dapat: %v", err)
	}

	if err := svc.UpdateMemberRole(ctx, 1, cohort.ID, 2, dto.UpdateMemberRoleRequest{Role: entity.RoleMentor}); err != nil {
		t.Fatalf("gagal mengubah role Bagas menjadi mentor: %v", err)
	}
	if err := svc.LeaveCohort(ctx, 1, cohort.ID); err != nil {
		t.Fatalf("Nadia harusnya sudah boleh keluar setelah ada mentor pengganti, dapat error: %v", err)
	}

	if _, err := svc.GetCohortDetail(ctx, 1, cohort.ID); err != response.ErrForbidden {
		t.Fatalf("mantan anggota harusnya ditolak akses (ErrForbidden), dapat: %v", err)
	}
}