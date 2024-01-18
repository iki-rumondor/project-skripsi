package domain

import "time"

type AssessmentSchedule struct {
	ID             uint   `gorm:"primaryKey"`
	Uuid           string `gorm:"not null, type:varchar(32)"`
	AssessmentID   uint
	ScheduleID     uint
	AssessmentType *AssessmentType `gorm:"foreignKey:AssessmentID"`
	Schedule       *Schedule
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
