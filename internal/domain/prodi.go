package domain

import (
	"time"

	"gorm.io/gorm"
)

type Prodi struct {
	ID        uint   `gorm:"primaryKey"`
	Nama      string `gorm:"not_null, varchar(20)"`
	Kaprodi   string `gorm:"not_null, varchar(20)"`
	JurusanID uint
	Jurusan   Jurusan
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Prodi) BeforeCreate(tx *gorm.DB) error {

	if err := tx.First(&Jurusan{}, "id = ?", p.JurusanID).Error; err != nil {
		return err
	}

	return nil
}
