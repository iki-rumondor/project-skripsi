package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rps struct {
	ID             uint   `gorm:"primaryKey"`
	Uuid           string `gorm:"not_null;unique;size:64"`
	SubjectID      uint   `gorm:"not_null"`
	AcademicYearID uint   `gorm:"not_null"`
	Note           *string
	Accept         bool `gorm:"not_null"`
	Status         bool `gorm:"not_null"`
	FileName       *string
	CreatedAt      int64 `gorm:"autoCreateTime:milli"`
	UpdatedAt      int64 `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	Subject        *Subject
	AcademicYear   *AcademicYear
}

func (m *Rps) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
