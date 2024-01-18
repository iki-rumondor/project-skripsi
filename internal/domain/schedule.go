package domain

import "time"

type Schedule struct {
	ID        uint   `gorm:"primaryKey"`
	Uuid      string `gorm:"not null, type:varchar(32)"`
	Start     string `gorm:"not null; type:date"`
	End       string `gorm:"not null; type:date"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
