package domain

import "time"

type AssessmentQuestion struct {
	ID               uint   `gorm:"primaryKey"`
	Uuid             string `gorm:"not null; varchar(32)"`
	Question         string `gorm:"not null; varchar(255)"`
	AssessmentTypeID uint
	AssessmentType   *AssessmentType
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
