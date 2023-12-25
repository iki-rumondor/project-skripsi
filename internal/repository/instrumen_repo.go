package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type InstrumenRepository interface {
	FindAllIndikator() (*[]domain.Indikator, error)
	FindIndikator(*domain.Indikator) (*domain.Indikator, error)
	CreateIndikator(*domain.Indikator) error
	UpdateIndikator(*domain.Indikator) error
	DeleteIndikator(*domain.Indikator) error
	CreateInstrumenType(*domain.InstrumenType) error
	FindAllInstrumenType() (*[]domain.InstrumenType, error) 
	DeleteInstrumenType(*domain.InstrumenType) (error)
	FindAllIndikatorType() (*[]domain.IndikatorType, error)
	CreateIndikatorType(data *domain.IndikatorType) error
	DeleteIndikatorType(data *domain.IndikatorType) error
}

