package entity

import "time"

type CohortRole string

const (
	RoleMentor CohortRole = "mentor"
	RoleMember CohortRole = "member"
)

type MembershipStatus string

const (
	MembershipActive MembershipStatus = "active"
	MembershipLeft   MembershipStatus = "left"
)

type InvitationStatus string

const (
	InvitationPending   InvitationStatus = "pending"
	InvitationAccepted  InvitationStatus = "accepted"
	InvitationDeclined  InvitationStatus = "declined"
	InvitationCancelled InvitationStatus = "cancelled"
)

type Cohort struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"size:100;not null"`
	Description string    `gorm:"type:text"`
	StartDate   time.Time `gorm:"type:date;not null"`
	EndDate     time.Time `gorm:"type:date;not null"`
	CreatedBy   uint      `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Membership struct {
	ID       uint             `gorm:"primaryKey"`
	CohortID uint             `gorm:"uniqueIndex:idx_cohort_user;not null"`
	UserID   uint             `gorm:"uniqueIndex:idx_cohort_user;not null"`
	Role     CohortRole       `gorm:"size:20;not null;default:'member'"`
	Status   MembershipStatus `gorm:"size:20;not null;default:'active'"`
	JoinedAt time.Time        `gorm:"not null"`
	LeftAt   *time.Time
	Cohort   Cohort `gorm:"foreignKey:CohortID;constraint:OnDelete:CASCADE"`
}

type Invitation struct {
	ID          uint             `gorm:"primaryKey"`
	CohortID    uint             `gorm:"index;not null"`
	InviterID   uint             `gorm:"not null"`
	Email       string           `gorm:"size:255;index;not null"`
	Role        CohortRole       `gorm:"size:20;not null;default:'member'"`
	Status      InvitationStatus `gorm:"size:20;not null;default:'pending'"`
	CreatedAt   time.Time
	RespondedAt *time.Time
	Cohort      Cohort `gorm:"foreignKey:CohortID;constraint:OnDelete:CASCADE"`
}