package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/cohort/contract"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/cohort/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/cohort/entity"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

type service struct {
	repo contract.CohortRepository
}

func New(repo contract.CohortRepository) contract.CohortService {
	return &service{repo: repo}
}

func (s *service) CreateCohort(ctx context.Context, userID uint, req dto.CreateCohortRequest) (*dto.CohortResponse, error) {
	startDate, err := parseDate(req.StartDate)
	if err != nil {
		return nil, response.ErrBadRequest
	}
	endDate, err := parseDate(req.EndDate)
	if err != nil || endDate.Before(startDate) {
		return nil, response.ErrBadRequest
	}

	now := time.Now()
	cohort := &entity.Cohort{
		Name:        strings.TrimSpace(req.Name),
		Description: strings.TrimSpace(req.Description),
		StartDate:   startDate,
		EndDate:     endDate,
		CreatedBy:   userID,
	}
	owner := &entity.Membership{
		UserID:   userID,
		Role:     entity.RoleMentor,
		Status:   entity.MembershipActive,
		JoinedAt: now,
	}

	if err := s.repo.CreateWithOwner(ctx, cohort, owner); err != nil {
		return nil, err
	}

	return &dto.CohortResponse{
		ID:          cohort.ID,
		Name:        cohort.Name,
		Description: cohort.Description,
		StartDate:   cohort.StartDate.Format("2006-01-02"),
		EndDate:     cohort.EndDate.Format("2006-01-02"),
		Role:        owner.Role,
		MemberCount: 1,
		JoinedAt:    owner.JoinedAt,
		CreatedAt:   cohort.CreatedAt,
	}, nil
}

func (s *service) ListMyCohorts(ctx context.Context, userID uint) ([]dto.CohortResponse, error) {
	cohorts, err := s.repo.FindUserCohorts(ctx, userID)
	if err != nil {
		return nil, err
	}
	if cohorts == nil {
		return []dto.CohortResponse{}, nil
	}
	return cohorts, nil
}

func (s *service) InviteMember(ctx context.Context, inviterID, cohortID uint, req dto.CreateInvitationRequest) (*dto.InvitationResponse, error) {
	inviterMember, err := s.repo.FindMembership(ctx, cohortID, inviterID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrForbidden
		}
		return nil, err
	}
	if inviterMember.Status != entity.MembershipActive || inviterMember.Role != entity.RoleMentor {
		return nil, response.ErrForbidden
	}

	cohort, err := s.repo.FindCohortByID(ctx, cohortID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, err
	}

	targetEmail := strings.ToLower(strings.TrimSpace(req.Email))
	targetUser, err := s.repo.FindUserByEmail(ctx, targetEmail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, err
	}

	existingMember, err := s.repo.FindMembership(ctx, cohortID, targetUser.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingMember != nil && existingMember.Status == entity.MembershipActive {
		return nil, response.ErrConflict
	}

	pendingInv, err := s.repo.FindPendingInvitation(ctx, cohortID, targetEmail)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if pendingInv != nil {
		return nil, response.ErrConflict
	}

	inviterUser, err := s.repo.FindUserByID(ctx, inviterID)
	if err != nil {
		return nil, err
	}

	inv := &entity.Invitation{
		CohortID:  cohortID,
		InviterID: inviterID,
		Email:     targetEmail,
		Role:      req.Role,
		Status:    entity.InvitationPending,
	}
	if err := s.repo.CreateInvitation(ctx, inv); err != nil {
		return nil, err
	}

	return &dto.InvitationResponse{
		ID:          inv.ID,
		CohortID:    cohort.ID,
		CohortName:  cohort.Name,
		InviterID:   inviterUser.ID,
		InviterName: inviterUser.Name,
		Email:       inv.Email,
		Role:        inv.Role,
		Status:      inv.Status,
		CreatedAt:   inv.CreatedAt,
	}, nil
}

func (s *service) ListCohortInvitations(ctx context.Context, userID, cohortID uint) ([]dto.InvitationResponse, error) {
	member, err := s.repo.FindMembership(ctx, cohortID, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrForbidden
		}
		return nil, err
	}
	if member.Status != entity.MembershipActive || member.Role != entity.RoleMentor {
		return nil, response.ErrForbidden
	}

	list, err := s.repo.FindCohortInvitations(ctx, cohortID)
	if err != nil {
		return nil, err
	}
	if list == nil {
		return []dto.InvitationResponse{}, nil
	}
	return list, nil
}

func (s *service) CancelInvitation(ctx context.Context, userID, cohortID, invitationID uint) error {
	member, err := s.repo.FindMembership(ctx, cohortID, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrForbidden
		}
		return err
	}
	if member.Status != entity.MembershipActive || member.Role != entity.RoleMentor {
		return response.ErrForbidden
	}

	inv, err := s.repo.FindInvitationByID(ctx, invitationID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrNotFound
		}
		return err
	}
	if inv.CohortID != cohortID {
		return response.ErrNotFound
	}
	if inv.Status != entity.InvitationPending {
		return response.ErrBadRequest
	}

	return s.repo.UpdateInvitationStatus(ctx, invitationID, entity.InvitationCancelled, nil)
}

func (s *service) ListMyInvitations(ctx context.Context, userID uint) ([]dto.InvitationResponse, error) {
	user, err := s.repo.FindUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, err
	}

	list, err := s.repo.FindUserPendingInvitations(ctx, user.Email)
	if err != nil {
		return nil, err
	}
	if list == nil {
		return []dto.InvitationResponse{}, nil
	}
	return list, nil
}

func (s *service) RespondInvitation(ctx context.Context, userID, invitationID uint, req dto.RespondInvitationRequest) error {
	user, err := s.repo.FindUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrNotFound
		}
		return err
	}

	inv, err := s.repo.FindInvitationByID(ctx, invitationID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrNotFound
		}
		return err
	}

	if !strings.EqualFold(inv.Email, user.Email) {
		return response.ErrForbidden
	}
	if inv.Status != entity.InvitationPending {
		return response.ErrBadRequest
	}

	now := time.Now()
	if req.Action == "decline" {
		return s.repo.UpdateInvitationStatus(ctx, inv.ID, entity.InvitationDeclined, &now)
	}

	existingMember, err := s.repo.FindMembership(ctx, inv.CohortID, userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if existingMember != nil && existingMember.Status == entity.MembershipActive {
		return response.ErrConflict
	}

	return s.repo.AcceptInvitationTx(ctx, inv, userID, now)
}

func parseDate(val string) (time.Time, error) {
	val = strings.TrimSpace(val)
	if t, err := time.Parse("2006-01-02", val); err == nil {
		return t, nil
	}
	return time.Parse(time.RFC3339, val)
}

func (s *service) GetCohortDetail(ctx context.Context, userID, cohortID uint) (*dto.CohortDetailResponse, error) {
	member, err := s.requireActiveMember(ctx, cohortID, userID)
	if err != nil {
		return nil, err
	}

	cohort, err := s.repo.FindCohortByID(ctx, cohortID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, err
	}

	members, err := s.repo.ListActiveMembers(ctx, cohortID)
	if err != nil {
		return nil, err
	}

	return &dto.CohortDetailResponse{
		ID:          cohort.ID,
		Name:        cohort.Name,
		Description: cohort.Description,
		StartDate:   cohort.StartDate.Format("2006-01-02"),
		EndDate:     cohort.EndDate.Format("2006-01-02"),
		MyRole:      member.Role,
		MemberCount: int64(len(members)),
		CreatedBy:   cohort.CreatedBy,
		CreatedAt:   cohort.CreatedAt,
		UpdatedAt:   cohort.UpdatedAt,
	}, nil
}

func (s *service) UpdateCohort(ctx context.Context, userID, cohortID uint, req dto.UpdateCohortRequest) (*dto.CohortDetailResponse, error) {
	member, err := s.requireActiveMentor(ctx, cohortID, userID)
	if err != nil {
		return nil, err
	}

	startDate, err := parseDate(req.StartDate)
	if err != nil {
		return nil, response.ErrBadRequest
	}
	endDate, err := parseDate(req.EndDate)
	if err != nil || endDate.Before(startDate) {
		return nil, response.ErrBadRequest
	}

	cohort, err := s.repo.FindCohortByID(ctx, cohortID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, err
	}

	cohort.Name = strings.TrimSpace(req.Name)
	cohort.Description = strings.TrimSpace(req.Description)
	cohort.StartDate = startDate
	cohort.EndDate = endDate

	if err := s.repo.UpdateCohort(ctx, cohort); err != nil {
		return nil, err
	}

	members, err := s.repo.ListActiveMembers(ctx, cohortID)
	if err != nil {
		return nil, err
	}

	return &dto.CohortDetailResponse{
		ID:          cohort.ID,
		Name:        cohort.Name,
		Description: cohort.Description,
		StartDate:   cohort.StartDate.Format("2006-01-02"),
		EndDate:     cohort.EndDate.Format("2006-01-02"),
		MyRole:      member.Role,
		MemberCount: int64(len(members)),
		CreatedBy:   cohort.CreatedBy,
		CreatedAt:   cohort.CreatedAt,
		UpdatedAt:   cohort.UpdatedAt,
	}, nil
}

func (s *service) DeleteCohort(ctx context.Context, userID, cohortID uint) error {
	if _, err := s.requireActiveMentor(ctx, cohortID, userID); err != nil {
		return err
	}

	if _, err := s.repo.FindCohortByID(ctx, cohortID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrNotFound
		}
		return err
	}

	return s.repo.DeleteCohort(ctx, cohortID)
}

func (s *service) ListMembers(ctx context.Context, userID, cohortID uint) ([]dto.MemberResponse, error) {
	if _, err := s.requireActiveMember(ctx, cohortID, userID); err != nil {
		return nil, err
	}

	members, err := s.repo.ListActiveMembers(ctx, cohortID)
	if err != nil {
		return nil, err
	}
	if members == nil {
		return []dto.MemberResponse{}, nil
	}
	return members, nil
}

func (s *service) UpdateMemberRole(ctx context.Context, actorID, cohortID, targetUserID uint, req dto.UpdateMemberRoleRequest) error {
	if _, err := s.requireActiveMentor(ctx, cohortID, actorID); err != nil {
		return err
	}

	target, err := s.repo.FindMembership(ctx, cohortID, targetUserID)
	if err != nil || target.Status != entity.MembershipActive {
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return response.ErrNotFound
	}

	if target.Role == req.Role {
		return nil
	}

	if target.Role == entity.RoleMentor && req.Role == entity.RoleMember {
		mentorCount, err := s.repo.CountActiveMentors(ctx, cohortID)
		if err != nil {
			return err
		}
		if mentorCount <= 1 {
			return response.ErrBadRequest
		}
	}

	return s.repo.UpdateMemberRole(ctx, cohortID, targetUserID, req.Role)
}

func (s *service) LeaveCohort(ctx context.Context, userID, cohortID uint) error {
	member, err := s.requireActiveMember(ctx, cohortID, userID)
	if err != nil {
		return err
	}

	if member.Role == entity.RoleMentor {
		mentorCount, err := s.repo.CountActiveMentors(ctx, cohortID)
		if err != nil {
			return err
		}
		if mentorCount <= 1 {
			return response.ErrBadRequest
		}
	}

	return s.repo.LeaveMembership(ctx, cohortID, userID, time.Now())
}

func (s *service) RemoveMember(ctx context.Context, actorID, cohortID, targetUserID uint) error {
	if _, err := s.requireActiveMentor(ctx, cohortID, actorID); err != nil {
		return err
	}

	target, err := s.repo.FindMembership(ctx, cohortID, targetUserID)
	if err != nil || target.Status != entity.MembershipActive {
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return response.ErrNotFound
	}

	if target.Role == entity.RoleMentor {
		mentorCount, err := s.repo.CountActiveMentors(ctx, cohortID)
		if err != nil {
			return err
		}
		if mentorCount <= 1 {
			return response.ErrBadRequest
		}
	}

	return s.repo.LeaveMembership(ctx, cohortID, targetUserID, time.Now())
}

func (s *service) requireActiveMember(ctx context.Context, cohortID, userID uint) (*entity.Membership, error) {
	m, err := s.repo.FindMembership(ctx, cohortID, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrForbidden
		}
		return nil, err
	}
	if m.Status != entity.MembershipActive {
		return nil, response.ErrForbidden
	}
	return m, nil
}

func (s *service) requireActiveMentor(ctx context.Context, cohortID, userID uint) (*entity.Membership, error) {
	m, err := s.requireActiveMember(ctx, cohortID, userID)
	if err != nil {
		return nil, err
	}
	if m.Role != entity.RoleMentor {
		return nil, response.ErrForbidden
	}
	return m, nil
}