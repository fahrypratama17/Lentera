package dto

import "github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/cohort/entity"

type CreateCohortRequest struct {
	Name        string `json:"name" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"max=1000"`
	StartDate   string `json:"start_date" binding:"required"` // format YYYY-MM-DD atau RFC3339
	EndDate     string `json:"end_date" binding:"required"`   // format YYYY-MM-DD atau RFC3339
}

type CreateInvitationRequest struct {
	Email string            `json:"email" binding:"required,email"`
	Role  entity.CohortRole `json:"role" binding:"required,oneof=mentor member"`
}

type RespondInvitationRequest struct {
	Action string `json:"action" binding:"required,oneof=accept decline"`
}

type UpdateCohortRequest struct {
	Name        string `json:"name" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"max=1000"`
	StartDate   string `json:"start_date" binding:"required"`
	EndDate     string `json:"end_date" binding:"required"`
}

type UpdateMemberRoleRequest struct {
	Role entity.CohortRole `json:"role" binding:"required,oneof=mentor member"`
}