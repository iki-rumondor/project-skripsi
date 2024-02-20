package models

import (
	"github.com/google/uuid"
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
	CreatedAt      int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt      int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	Teacher        *Teacher
	Subject        *Subject
	AcademicYear   *AcademicYear
}

func (m *TeacherAttendence) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
