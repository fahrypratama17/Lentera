package entity

import "time"

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
