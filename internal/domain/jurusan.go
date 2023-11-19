package domain

import "time"

type Jurusan struct {
	ID        uint   `gorm:"primaryKey"`
	Nama      string `gorm:"not_null, varchar(20)"`
	CreatedAt time.Time
}
