package entity

import "time"

type Progress struct {
	ID         uint   `gorm:"primaryKey"`
	CohortID   uint   `gorm:"not null;index"`
	MaterialID uint   `gorm:"not null;uniqueIndex:uq_progress_material_user"`
	UserID     uint   `gorm:"not null;uniqueIndex:uq_progress_material_user;index"`
	Status     string `gorm:"size:20;not null"`

	// pengumpulan terbaru (ditulis saat submit)
	SubmissionContent *string `gorm:"type:text"`
	SubmissionURL     *string `gorm:"size:500"`
	SubmittedAt       *time.Time

	// keputusan review terakhir (ditulis saat review)
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

func (Progress) TableName() string { return "progress" }
