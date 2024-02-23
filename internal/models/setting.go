package models

type Setting struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"not_null;size:64;unique"`
	Value string `gorm:"not_null;size:64"`
}
