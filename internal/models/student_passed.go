package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StudentPassed struct {
	ID             uint   `gorm:"primaryKey"`
	Uuid           string `gorm:"not_null,unique;size:64"`
	SubjectID      uint   `gorm:"not_null"`
	AcademicYearID uint   `gorm:"not_null"`
	StudentAmount  uint   `gorm:"not_null"`
	PassedAmount   uint   `gorm:"not_null;"`
	CreatedAt      int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt      int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	Subject        *Subject
	AcademicYear   *AcademicYear
}

func (m *StudentPassed) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
