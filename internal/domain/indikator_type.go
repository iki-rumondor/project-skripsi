package domain

import (
	"time"

	"gorm.io/gorm"
)

type IndikatorType struct {
	ID          uint   `gorm:"primaryKey"`
	Description string `gorm:"not null; varchar(120)"`
	TypeID      uint

	InstrumenType InstrumenType `gorm:"foreignKey:TypeID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *IndikatorType) BeforeCreate(tx *gorm.DB) error {

	if err := tx.First(&InstrumenType{}, "id = ?", m.TypeID).Error; err != nil {
		return err;
	}

	return nil
}

func (m *IndikatorType) BeforeDelete(tx *gorm.DB) error {

	if err := tx.First(&IndikatorType{}, "id = ?", m.ID).Error; err != nil {
		return err;
	}

	return nil
}