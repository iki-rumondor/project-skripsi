package domain

import (
	"errors"
	"time"

	"github.com/iki-rumondor/init-golang-service/internal/utils"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not_null;varchar(120)"`
	Password string `gorm:"not_null;varchar(120)"`
	Role     string `gorm:"not_null;varchar(16)"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) error {

	var user User

	if result := tx.First(&user, "username = ? AND id != ?", u.Username, u.ID).RowsAffected; result > 0 {
		return errors.New("the username has already been taken")
	}

	hashPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashPass
	return nil
}

type Role struct {
	ID   uint
	Name string
}
