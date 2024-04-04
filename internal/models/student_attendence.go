package models

import (
	"github.com/google/uuid"
	"github.com/iki-rumondor/go-monev/internal/http/response"
	"gorm.io/gorm"
)

type StudentAttendence struct {
	ID             uint   `gorm:"primaryKey"`
	Uuid           string `gorm:"not_null,unique;size:64"`
	SubjectID      uint   `gorm:"not_null"`
	Class          string `gorm:"not_null;size:2"`
	AcademicYearID uint   `gorm:"not_null"`
	StudentAmount  uint   `gorm:"not_null"`
	PassedAmount   uint   `gorm:"not_null"`
	FinalAmount    uint   `gorm:"not_null"`
	Middle         uint   `gorm:"not_null;"`
	Last           uint   `gorm:"not_null;"`
	CreatedAt      int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt      int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	Subject        *Subject
	AcademicYear   *AcademicYear
}

func (m *StudentAttendence) BeforeCreate(tx *gorm.DB) error {
	if result := tx.First(&StudentAttendence{}, "class = ? AND subject_id = ?", m.Class, m.SubjectID).RowsAffected; result > 0 {
		return response.BADREQ_ERR("Kelas Sudah Didaftarkan")
	}
	m.Uuid = uuid.NewString()
	return nil
}
