package domain

import "time"

type Response struct {
	ID         uint   `gorm:"primaryKey"`
	Uuid       string `gorm:"not null; varchar(32)"`
	UserID     uint
	QuestionID uint
	Response   bool
	User       *User
	Question   *AssessmentQuestion `gorm:"foreignKey:QuestionID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
