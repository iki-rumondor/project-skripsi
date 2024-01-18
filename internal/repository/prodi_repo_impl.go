package repository

import (
	"fmt"

	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type ProdiRepoImplementation struct {
	db *gorm.DB
}

func NewProdiRepository(db *gorm.DB) ProdiRepository {
	return &ProdiRepoImplementation{
		db: db,
	}
}

func (r *ProdiRepoImplementation) FindAllProdi() (*[]domain.Prodi, error) {
	var prodi []domain.Prodi
	if err := r.db.Preload("Jurusan").Find(&prodi).Error; err != nil {
		return nil, err
	}

	return &prodi, nil
}

func (r *ProdiRepoImplementation) FindProdi(id uint) (*domain.Prodi, error) {
	var prodi domain.Prodi
	if err := r.db.Preload("Jurusan").First(&prodi, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &prodi, nil
}

func (r *ProdiRepoImplementation) FindProdiByUuid(uuid string) (*domain.Prodi, error) {
	var prodi domain.Prodi
	if err := r.db.Preload("Jurusan").Preload("Subject").First(&prodi, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &prodi, nil
}

func (r *ProdiRepoImplementation) FindBy(column string, value interface{}) (*domain.Prodi, error) {
	var prodi domain.Prodi
	if err := r.db.Preload("Jurusan").Preload("Subject").First(&prodi, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &prodi, nil
}

func (r *ProdiRepoImplementation) CreateProdi(prodi *domain.Prodi) error {
	return r.db.Save(&prodi).Error
}

func (r *ProdiRepoImplementation) UpdateProdi(prodi *domain.Prodi) error {
	return r.db.Model(&prodi).Updates(&prodi).Error
}

func (r *ProdiRepoImplementation) DeleteProdi(prodi *domain.Prodi) error {
	return r.db.Delete(&prodi).Error
}
