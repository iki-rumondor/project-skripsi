package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AcademicYear struct {
	ID        uint   `gorm:"primaryKey"`
	Uuid      string `gorm:"not_null,unique,type:varchar,size:64"`
	Name      string `gorm:"not_null,type:varchar,size:16"`
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"autoCreateTime:milli,autoUpdateTime:milli"`
}

func (m *AcademicYear) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
