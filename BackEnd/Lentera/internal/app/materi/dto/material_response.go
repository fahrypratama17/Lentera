package dto

import "time"

type MaterialResponse struct {
	ID             uint              `json:"id"`
	CohortID       uint              `json:"cohort_id"`
	Title          string            `json:"title"`
	Guide          string            `json:"guide"`
	References     []string          `json:"references"`
	Deadline       *time.Time        `json:"deadline"`
	IsPastDeadline bool              `json:"is_past_deadline"`
	Position       int               `json:"position"`
	CreatedAt      time.Time         `json:"created_at"`
	UpdatedAt      time.Time         `json:"updated_at"`
	MyProgress     *ProgressResponse `json:"my_progress"`
}

type ProgressResponse struct {
	ID          *uint               `json:"id"`
	UserID      uint                `json:"user_id"`
	MaterialID  uint                `json:"material_id"`
	Status      string              `json:"status"`
	IsLate      bool                `json:"is_late"`
	StartedAt   *time.Time          `json:"started_at"`
	CompletedAt *time.Time          `json:"completed_at"`
	Submission  *SubmissionResponse `json:"submission"`
	LastReview  *ReviewResponse     `json:"last_review"`
	UpdatedAt   *time.Time          `json:"updated_at"`
}

type SubmissionResponse struct {
	Content     string    `json:"content"`
	WorkURL     *string   `json:"work_url"`
	SubmittedAt time.Time `json:"submitted_at"`
	IsLate      bool      `json:"is_late"`
}

type ReviewResponse struct {
	Decision   string    `json:"decision"`
	Note       *string   `json:"note"`
	Reviewer   UserBrief `json:"reviewer"`
	ReviewedAt time.Time `json:"reviewed_at"`
}

type UserBrief struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type PersonalSummary struct {
	TotalMaterials int     `json:"total_materials"`
	NotStarted     int     `json:"not_started"`
	InProgress     int     `json:"in_progress"`
	Submitted      int     `json:"submitted"`
	Completed      int     `json:"completed"`
	Late           int     `json:"late"`
	CompletionRate float64 `json:"completion_rate"`
}
