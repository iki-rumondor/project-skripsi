package domain

import "time"

type Subject struct {
	ID        uint   `gorm:"primaryKey"`
	Uuid      string `gorm:"not null, type:varchar(32)"`
	Name      string `gorm:"not null, type:varchar(64)"`
	Code      string `gorm:"not null, type:varchar(16)"`
	ProdiID   uint   `gorm:"not null; type:int"`
	Prodi     *Prodi
	CreatedAt time.Time
	UpdatedAt time.Time
}
