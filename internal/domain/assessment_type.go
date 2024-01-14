package domain

import "time"

type AssessmentType struct {
	ID        uint   `gorm:"primaryKey"`
	Uuid      string `gorm:"not null; varchar(32)"`
	Type      string `gorm:"not null; varchar(64)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
