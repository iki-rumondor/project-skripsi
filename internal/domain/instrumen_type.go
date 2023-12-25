package domain

import (
	"time"

	"gorm.io/gorm"
)

type InstrumenType struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null; varchar(120)"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *InstrumenType) BeforeUpdate(tx *gorm.DB) error {

	if err := tx.First(&InstrumenType{}, "id = ?", m.ID).Error; err != nil {
		return err
	}

	return nil
}