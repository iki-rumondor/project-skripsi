package repository

import (
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type UserRepoImplementation struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepoImplementation{
		db: db,
	}
}

func (r *UserRepoImplementation) CreateUser(model *domain.User) error {
	return r.db.Create(&model).Error
}

func (r *UserRepoImplementation) FindAllUser() (*[]domain.User, error) {
	var result []domain.User
	if err := r.db.Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *UserRepoImplementation) FindUserByUuid(uuid string) (*domain.User, error) {
	var result domain.User
	if err := r.db.First(&result, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *UserRepoImplementation) UpdateUser(model *domain.User) error {
	return r.db.Model(&model).Updates(&model).Error
}

func (r *UserRepoImplementation) DeleteUser(model *domain.User) error {
	return r.db.Delete(&model).Error
}
