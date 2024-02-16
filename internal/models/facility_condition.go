package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FacilityCondition struct {
	ID             uint   `gorm:"primaryKey"`
	Uuid           string `gorm:"not_null;unique;size:64"`
	FacilityID     uint   `gorm:"not_null"`
	AcademicYearID uint   `gorm:"not_null"`
	Amount         uint   `gorm:"not_null;size:3"`
	Unit           string `gorm:"not_null;size:32"`
	Deactive       uint   `gorm:"not_null;size:3"`
	Note           string `gorm:"not_null;size:64"`
	CreatedAt      int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt      int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	Facility       *Facility
	AcademicYear   *AcademicYear
}

func (m *FacilityCondition) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
