package repository

import (
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type AuthRepoImplementation struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepoImplementation{
		db: db,
	}
}

func (r *AuthRepoImplementation) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepoImplementation) FindUsers() (*[]domain.User, error) {
	var users []domain.User
	if err := r.db.Preload("Role").Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (r *AuthRepoImplementation) SaveUser(user *domain.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}

	return nil
}
