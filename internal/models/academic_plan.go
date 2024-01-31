package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AcademicPlan struct {
	ID             uint   `gorm:"primaryKey"`
	Uuid           string `gorm:"not_null;unique;size:64"`
	SubjectID      uint   `gorm:"not_null;size:2"`
	AcademicYearID uint   `gorm:"not_null;size:2"`
	Name           string `gorm:"not_null;size:32"`
	Available      bool   `gorm:"not_null"`
	Note           string `gorm:"not_null;size:144"`
	CreatedAt      int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt      int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	Subject        *Subject
	AcademicYear   *AcademicYear
}

func (m *AcademicPlan) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
