package repository

import (
	"fmt"

	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepoInterface {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindUserBy(column string, value interface{}) (*models.User, error) {
	var user models.User
	if err := r.db.First(fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
