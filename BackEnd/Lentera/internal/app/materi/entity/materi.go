package entity

import "time"

const (
	StatusNotStarted = "not_started"
	StatusInProgress = "in_progress"
	StatusSubmitted  = "submitted"
	StatusCompleted  = "completed"

	RoleMentor = "mentor"
	RoleMember = "member"
)

type Material struct {
	ID         uint     `gorm:"primaryKey"`
	CohortID   uint     `gorm:"not null;index:idx_materials_cohort_position,priority:1"`
	Title      string   `gorm:"size:200;not null"`
	Guide      string   `gorm:"type:text;not null"`
	References []string `gorm:"type:jsonb;serializer:json"`
	Deadline   *time.Time
	Position   int `gorm:"not null;index:idx_materials_cohort_position,priority:2"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Progress struct {
	ID         uint   `gorm:"primaryKey"`
	CohortID   uint   `gorm:"not null;index"`
	MaterialID uint   `gorm:"not null;uniqueIndex:uq_progress_material_user"`
	UserID     uint   `gorm:"not null;uniqueIndex:uq_progress_material_user;index"`
	Status     string `gorm:"size:20;not null"`

	SubmissionContent *string `gorm:"type:text"`
	SubmissionURL     *string `gorm:"size:500"`
	SubmittedAt       *time.Time

	ReviewDecision *string `gorm:"size:30"`
	ReviewNote     *string `gorm:"type:text"`
	ReviewerID     *uint
	ReviewedAt     *time.Time

	StartedAt   *time.Time
	CompletedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Material *Material `gorm:"foreignKey:MaterialID;constraint:OnDelete:CASCADE"`
}

// MemberInfo = data anggota hasil join memberships + users.
// BUKAN tabel (tidak di-AutoMigrate), hanya wadah hasil query.
type MemberInfo struct {
	UserID uint
	Name   string
	Role   string // mentor | member
	Status string // active | left
}

func (Progress) TableName() string { return "progress" }
