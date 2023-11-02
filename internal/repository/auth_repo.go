package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type AuthRepository interface {
	FindByEmail(string) (*domain.User, error)
	SaveUser(*domain.User) error
	FindUsers() (*[]domain.User, error)
}
