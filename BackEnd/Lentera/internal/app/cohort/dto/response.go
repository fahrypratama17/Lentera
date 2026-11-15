package dto

import (
	"time"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/cohort/entity"
)

type CohortResponse struct {
	ID          uint              `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	StartDate   string            `json:"start_date"`
	EndDate     string            `json:"end_date"`
	Role        entity.CohortRole `json:"role"`
	MemberCount int64             `json:"member_count"`
	JoinedAt    time.Time         `json:"joined_at"`
	CreatedAt   time.Time         `json:"created_at"`
}

type InvitationResponse struct {
	ID          uint                    `json:"id"`
	CohortID    uint                    `json:"cohort_id"`
	CohortName  string                  `json:"cohort_name"`
	InviterID   uint                    `json:"inviter_id"`
	InviterName string                  `json:"inviter_name"`
	Email       string                  `json:"email"`
	Role        entity.CohortRole       `json:"role"`
	Status      entity.InvitationStatus `json:"status"`
	CreatedAt   time.Time               `json:"created_at"`
	RespondedAt *time.Time              `json:"responded_at,omitempty"`
}

type CohortDetailResponse struct {
	ID          uint              `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	StartDate   string            `json:"start_date"`
	EndDate     string            `json:"end_date"`
	MyRole      entity.CohortRole `json:"my_role"`
	MemberCount int64             `json:"member_count"`
	CreatedBy   uint              `json:"created_by"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

type MemberResponse struct {
	UserID   uint              `json:"user_id"`
	Name     string            `json:"name"`
	Email    string            `json:"email"`
	Role     entity.CohortRole `json:"role"`
	JoinedAt time.Time         `json:"joined_at"`
}