package domain

import (
	"time"

	"gorm.io/gorm"
)

type Indikator struct {
	ID          uint   `gorm:"primaryKey"`
	Description string `gorm:"not null; varchar(120)"`
	IndikatorTypeID      uint

	IndikatorType IndikatorType

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Indikator) BeforeCreate(tx *gorm.DB) error {

	if err := tx.First(&IndikatorType{}, "id = ?", m.IndikatorTypeID).Error; err != nil {
		return err
	}

	return nil
}

func (m *Indikator) BeforeUpdate(tx *gorm.DB) error {

	if err := tx.First(&Indikator{}, "id = ?", m.ID).Error; err != nil {
		return err
	}

	return nil
}

func (m *Indikator) BeforeDelete(tx *gorm.DB) error {

	if err := tx.First(&Indikator{}, "id = ?", m.ID).Error; err != nil {
		return err
	}

	return nil
}