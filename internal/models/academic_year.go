package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AcademicYear struct {
	ID               uint   `gorm:"primaryKey"`
	Uuid             string `gorm:"not_null;unique;size:64"`
	Year             string `gorm:"not_null;size:16"`
	Semester         string `gorm:"not_null;size:16"`
	FirstDate        string `gorm:"not_null;size:32"`
	FirstDays        string `gorm:"not_null;size:32"`
	MiddleDate       string `gorm:"not_null;size:32"`
	MiddleDays       string `gorm:"not_null;size:32"`
	MiddleLastDate   string `gorm:"not_null;size:32"`
	MiddleLastDays   string `gorm:"not_null;size:32"`
	LastDate         string `gorm:"not_null;size:32"`
	LastDays         string `gorm:"not_null;size:32"`
	CreatedAt        int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt        int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	AcademicPlans    *[]AcademicPlan
	PracticalTools   *[]PracticalTool
	PracticalModules *[]PracticalModule
}

func (m *AcademicYear) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
