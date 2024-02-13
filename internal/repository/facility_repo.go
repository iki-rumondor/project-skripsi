package repository

import (
	"fmt"

	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
)

type FacilityRepository struct {
	db *gorm.DB
}

func NewFacilityRepository(db *gorm.DB) interfaces.FacilityRepoInterface {
	return &FacilityRepository{
		db: db,
	}
}

func (r *FacilityRepository) FindFacilities(userUuid string) (*[]models.Facility, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var result []models.Facility
	if err := r.db.Preload("Department").Find(&result, "department_id = ?", user.Department.ID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *FacilityRepository) FindUserFacility(userUuid, uuid string) (*models.Facility, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var result models.Facility
	if err := r.db.Preload("Department").First(&result, "uuid = ? AND department_id = ?", uuid, user.Department.ID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *FacilityRepository) FindUserBy(column string, value interface{}) (*models.User, error) {
	var result models.User
	if err := r.db.Preload("Department").First(&result, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *FacilityRepository) CreateFacility(model *models.Facility) error {
	return r.db.Create(model).Error
}

func (r *FacilityRepository) UpdateFacility(model *models.Facility) error {
	return r.db.Updates(model).Error
}

func (r *FacilityRepository) DeleteFacility(model *models.Facility) error {
	return r.db.Delete(model).Error
}
