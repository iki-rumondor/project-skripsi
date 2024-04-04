package models

import (
	"github.com/google/uuid"
	"github.com/iki-rumondor/go-monev/internal/http/response"
	"gorm.io/gorm"
)

type TeacherAttendence struct {
	ID             uint   `gorm:"primaryKey"`
	Uuid           string `gorm:"not_null,unique;size:64"`
	SubjectID      uint   `gorm:"not_null"`
	AcademicYearID uint   `gorm:"not_null"`
	TeacherID      uint   `gorm:"not_null"`
	Middle         uint   `gorm:"not_null;default:0"`
	Last           uint   `gorm:"not_null;default:0"`
	GradeOnTime    bool   `gorm:"not_null"`
	Class          string `gorm:"not_null;size:2"`
	CreatedAt      int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt      int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	Teacher        *Teacher
	Subject        *Subject
	AcademicYear   *AcademicYear
}

func (m *TeacherAttendence) BeforeCreate(tx *gorm.DB) error {
	if result := tx.First(&TeacherAttendence{}, "class = ? AND subject_id = ?", m.Class, m.SubjectID).RowsAffected; result > 0 {
		return response.BADREQ_ERR("Kelas Sudah Didaftarkan")
	}
	m.Uuid = uuid.NewString()
	return nil
}
