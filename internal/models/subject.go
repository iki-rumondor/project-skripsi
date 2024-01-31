package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Subject struct {
	ID            uint   `gorm:"primaryKey"`
	Uuid          string `gorm:"not_null;unique;size:64"`
	Name          string `gorm:"not_null;size:32"`
	Code          string `gorm:"not_null;unique;size:16"`
	DepartmentID  uint   `gorm:"not_null"`
	CreatedAt     int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt     int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	Department    *Department
	AcademicPlans *[]AcademicPlan
}

func (m *Subject) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
