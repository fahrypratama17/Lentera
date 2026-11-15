package contract

import (
	"context"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/cohort/dto"
)

type CohortService interface {
	CreateCohort(ctx context.Context, userID uint, req dto.CreateCohortRequest) (*dto.CohortResponse, error)
	ListMyCohorts(ctx context.Context, userID uint) ([]dto.CohortResponse, error)

	InviteMember(ctx context.Context, inviterID, cohortID uint, req dto.CreateInvitationRequest) (*dto.InvitationResponse, error)
	ListCohortInvitations(ctx context.Context, userID, cohortID uint) ([]dto.InvitationResponse, error)
	CancelInvitation(ctx context.Context, userID, cohortID, invitationID uint) error

	ListMyInvitations(ctx context.Context, userID uint) ([]dto.InvitationResponse, error)
	RespondInvitation(ctx context.Context, userID, invitationID uint, req dto.RespondInvitationRequest) error

	GetCohortDetail(ctx context.Context, userID, cohortID uint) (*dto.CohortDetailResponse, error)
	UpdateCohort(ctx context.Context, userID, cohortID uint, req dto.UpdateCohortRequest) (*dto.CohortDetailResponse, error)
	DeleteCohort(ctx context.Context, userID, cohortID uint) error

	ListMembers(ctx context.Context, userID, cohortID uint) ([]dto.MemberResponse, error)
	UpdateMemberRole(ctx context.Context, actorID, cohortID, targetUserID uint, req dto.UpdateMemberRoleRequest) error
	LeaveCohort(ctx context.Context, userID, cohortID uint) error
	RemoveMember(ctx context.Context, actorID, cohortID, targetUserID uint) error
}