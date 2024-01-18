package domain

import (
	"time"

	"gorm.io/gorm"
)

type Prodi struct {
	ID         uint   `gorm:"primaryKey"`
	Uuid       string `gorm:"not_null, type:varchar(32)"`
	Nama       string `gorm:"not_null, type:varchar(20)"`
	Kaprodi    string `gorm:"not_null, type:varchar(20)"`
	Credential string `gorm:"not_null, type:varchar(64)"`
	JurusanID  uint
	Jurusan    *Jurusan
	Subject    *[]Subject
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (p *Prodi) BeforeCreate(tx *gorm.DB) error {

	if err := tx.First(&Jurusan{}, "id = ?", p.JurusanID).Error; err != nil {
		return err
	}

	return nil
}
