package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PracticalModule struct {
	ID             uint   `gorm:"primaryKey"`
	Uuid           string `gorm:"not_null,unique,type:varchar,size:64"`
	SubjectID      uint   `gorm:"not_null,type:uint,size:11"`
	LaboratoryID   uint   `gorm:"not_null,type:uint,size:11"`
	AcademicYearID uint   `gorm:"not_null,type:uint,size:11"`
	Available      bool   `gorm:"not_null,type:boolean"`
	Note           string `gorm:"not_null,type:varchar,size:144"`
	CreatedAt      int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt      int64  `gorm:"autoCreateTime:milli,autoUpdateTime:milli"`
	Subject        *Subject
	Laboratory     *Laboratory
	AcademicYear   *AcademicYear
}

func (m *PracticalModule) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
