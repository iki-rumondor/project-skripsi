package domain

import (
	"errors"
	"time"

	"github.com/iki-rumondor/init-golang-service/internal/utils"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Uuid     string `gorm:"not_null;varchar(120)"`
	Username string `gorm:"unique;not_null;varchar(120)"`
	Email    string `gorm:"unique;not_null;varchar(120)"`
	Password string `gorm:"not_null;varchar(120)"`
	RoleID   uint
	Role     Role

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeSave(tx *gorm.DB) error {

	var user User
	if result := tx.First(&user, "email = ? AND id != ?", u.Email, u.ID).RowsAffected; result > 0 {
		return errors.New("the email has already been taken")
	}

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
