package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type UserRepository interface {
	CreateUser(model *domain.User) error
	FindAllUser() (*[]domain.User, error)
	FindUserByUuid(uuid string) (*domain.User, error)
	UpdateUser(model *domain.User) error
	DeleteUser(model *domain.User) error
}
