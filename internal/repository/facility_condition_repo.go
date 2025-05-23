package repository

import (
	"fmt"

	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FacilityConditionRepository struct {
	db *gorm.DB
}

func NewFacilityConditionRepository(db *gorm.DB) interfaces.FacilityConditionRepoInterface {
	return &FacilityConditionRepository{
		db: db,
	}
}

func (r *FacilityConditionRepository) FindFacilityConditionsByYear(userUuid string, yearID uint) (*[]models.FacilityCondition, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var result []models.FacilityCondition

	if err := r.db.Joins("Facility").Preload("AcademicYear").Find(&result, "Facility.department_id = ? AND Academic_year_id = ?", user.Department.ID, yearID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *FacilityConditionRepository) FindFacilityOptions(userUuid string, yearID uint) (*[]models.Facility, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	subQuery := r.db.Model(&models.FacilityCondition{}).Where("academic_year_id = ?", yearID).Select("facility_id")

	var result []models.Facility

	if err := r.db.Find(&result, "department_id = ? AND id NOT IN (?)", user.Department.ID, subQuery).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *FacilityConditionRepository) FindUserFacilityCondition(userUuid, uuid string) (*models.FacilityCondition, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var model models.FacilityCondition
	if err := r.db.Joins("Facility").Preload("AcademicYear").First(&model, "Facility_conditions.uuid = ? AND Facility.department_id = ?", uuid, user.Department.ID).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *FacilityConditionRepository) FindFacilityBy(column string, value interface{}) (*models.Facility, error) {
	var model models.Facility
	if err := r.db.Preload("Department.User").First(&model, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *FacilityConditionRepository) FindAcademicYearBy(column string, value interface{}) (*models.AcademicYear, error) {
	var model models.AcademicYear
	if err := r.db.First(&model, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *FacilityConditionRepository) CreateFacilityCondition(model *models.FacilityCondition) error {
	return r.db.Create(model).Error
}

func (r *FacilityConditionRepository) UpdateFacilityCondition(model *models.FacilityCondition) error {
	return r.db.Updates(model).Error
}

func (r *FacilityConditionRepository) DeleteFacilityCondition(model *models.FacilityCondition) error {
	return r.db.Delete(model).Error
}

func (r *FacilityConditionRepository) FindDepartment(uuid string) (*models.Department, error) {
	var department models.Department
	if err := r.db.First(&department, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &department, nil
}

func (r *FacilityConditionRepository) FindByDepartment(departmentID, yearID uint) (*[]models.FacilityCondition, error) {
	var result []models.FacilityCondition

	if err := r.db.Joins("Facility").Preload(clause.Associations).Find(&result, "Facility.department_id = ? AND academic_year_id = ?", departmentID, yearID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
