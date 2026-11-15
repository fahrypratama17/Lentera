package contract

import (
	"context"
	"time"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/cohort/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/cohort/entity"
)

type UserBasic struct {
	ID    uint
	Name  string
	Email string
}

type CohortRepository interface {
	CreateWithOwner(ctx context.Context, cohort *entity.Cohort, owner *entity.Membership) error
	FindUserCohorts(ctx context.Context, userID uint) ([]dto.CohortResponse, error)
	FindCohortByID(ctx context.Context, cohortID uint) (*entity.Cohort, error)
	FindMembership(ctx context.Context, cohortID, userID uint) (*entity.Membership, error)

	FindUserByID(ctx context.Context, userID uint) (*UserBasic, error)
	FindUserByEmail(ctx context.Context, email string) (*UserBasic, error)

	CreateInvitation(ctx context.Context, inv *entity.Invitation) error
	FindPendingInvitation(ctx context.Context, cohortID uint, email string) (*entity.Invitation, error)
	FindInvitationByID(ctx context.Context, invitationID uint) (*entity.Invitation, error)
	FindUserPendingInvitations(ctx context.Context, email string) ([]dto.InvitationResponse, error)
	FindCohortInvitations(ctx context.Context, cohortID uint) ([]dto.InvitationResponse, error)
	UpdateInvitationStatus(ctx context.Context, invitationID uint, status entity.InvitationStatus, respondedAt *time.Time) error
	AcceptInvitationTx(ctx context.Context, inv *entity.Invitation, userID uint, now time.Time) error

	UpdateCohort(ctx context.Context, cohort *entity.Cohort) error
	DeleteCohort(ctx context.Context, cohortID uint) error
	ListActiveMembers(ctx context.Context, cohortID uint) ([]dto.MemberResponse, error)
	CountActiveMentors(ctx context.Context, cohortID uint) (int64, error)
	UpdateMemberRole(ctx context.Context, cohortID, userID uint, role entity.CohortRole) error
	LeaveMembership(ctx context.Context, cohortID, userID uint, leftAt time.Time) error
}