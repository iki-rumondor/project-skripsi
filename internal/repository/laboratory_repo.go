package repository

import (
	"fmt"

	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
)

type LaboratoryRepository struct {
	db *gorm.DB
}

func NewLaboratoryRepository(db *gorm.DB) interfaces.LaboratoryRepoInterface {
	return &LaboratoryRepository{
		db: db,
	}
}

func (r *LaboratoryRepository) FindLaboratories(userUuid string) (*[]models.Laboratory, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var Laboratorys []models.Laboratory
	if err := r.db.Preload("Department").Find(&Laboratorys, "department_id = ?", user.Department.ID).Error; err != nil {
		return nil, err
	}

	return &Laboratorys, nil
}

func (r *LaboratoryRepository) FindUserLaboratory(userUuid, uuid string) (*models.Laboratory, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var Laboratory models.Laboratory
	if err := r.db.Preload("Department").First(&Laboratory, "uuid = ? AND department_id = ?", uuid, user.Department.ID).Error; err != nil {
		return nil, err
	}

	return &Laboratory, nil
}

func (r *LaboratoryRepository) FindUserBy(column string, value interface{}) (*models.User, error) {
	var result models.User
	if err := r.db.Preload("Department").First(&result, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *LaboratoryRepository) CreateLaboratory(Laboratory *models.Laboratory) error {
	return r.db.Create(Laboratory).Error
}

func (r *LaboratoryRepository) UpdateLaboratory(Laboratory *models.Laboratory) error {
	if err := r.db.First(&models.Laboratory{}, "id = ? AND department_id = ?", Laboratory.ID, Laboratory.DepartmentID).Error; err != nil {
		return err
	}

	return r.db.Updates(Laboratory).Error
}

func (r *LaboratoryRepository) DeleteLaboratory(Laboratory *models.Laboratory) error {
	if err := r.db.First(&models.Laboratory{}, "id = ? AND department_id = ?", Laboratory.ID, Laboratory.DepartmentID).Error; err != nil {
		return err
	}

	return r.db.Delete(Laboratory).Error
}
