package repository

import (
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type InstrumenRepoImplementation struct {
	db *gorm.DB
}

func NewInstrumenRepository(db *gorm.DB) InstrumenRepository {
	return &InstrumenRepoImplementation{
		db: db,
	}
}

func (r *InstrumenRepoImplementation) FindAllIndikator() (*[]domain.Indikator, error) {
	var res []domain.Indikator
	if err := r.db.Preload("IndikatorType").Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *InstrumenRepoImplementation) FindIndikator(data *domain.Indikator) (*domain.Indikator, error) {
	var res domain.Indikator
	if err := r.db.Preload("IndikatorType").First(&res, "id = ?", data.ID).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *InstrumenRepoImplementation) CreateIndikator(data *domain.Indikator) error {
	return r.db.Save(data).Error
}

func (r *InstrumenRepoImplementation) UpdateIndikator(data *domain.Indikator) error {
	return r.db.Model(data).Where("id =  ?", data.ID).Updates(data).Error
}

func (r *InstrumenRepoImplementation) DeleteIndikator(data *domain.Indikator) error {
	return r.db.Delete(data).Error
}

func (r *InstrumenRepoImplementation) CreateInstrumenType(data *domain.InstrumenType) error {
	return r.db.Save(data).Error
}

func (r *InstrumenRepoImplementation) FindAllInstrumenType() (*[]domain.InstrumenType, error) {
	var res []domain.InstrumenType
	if err := r.db.Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *InstrumenRepoImplementation) DeleteInstrumenType(data *domain.InstrumenType) error {
	return r.db.Delete(data).Error
}

func (r *InstrumenRepoImplementation) FindAllIndikatorType() (*[]domain.IndikatorType, error) {
	var res []domain.IndikatorType
	if err := r.db.Preload("InstrumenType").Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *InstrumenRepoImplementation) CreateIndikatorType(data *domain.IndikatorType) error {
	return r.db.Save(data).Error
}

func (r *InstrumenRepoImplementation) DeleteIndikatorType(data *domain.IndikatorType) error {
	return r.db.Delete(data).Error
}
