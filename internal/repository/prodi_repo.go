package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type ProdiRepository interface {
	FindAllProdi() (*[]domain.Prodi, error)
	CreateProdi(*domain.Prodi) error
	FindProdi(uint) (*domain.Prodi, error)
	FindProdiByUuid(string) (*domain.Prodi, error)
	UpdateProdi(*domain.Prodi) error
	DeleteProdi(*domain.Prodi) error
}
