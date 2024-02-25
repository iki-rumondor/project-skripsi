package repository

import (
	"fmt"

	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
)

type AcademicPlanRepository struct {
	db *gorm.DB
}

func NewAcademicPlanRepository(db *gorm.DB) interfaces.AcademicPlanRepoInterface {
	return &AcademicPlanRepository{
		db: db,
	}
}

func (r *AcademicPlanRepository) FindAcademicPlans(userUuid string, yearID uint) (*[]models.AcademicPlan, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var result []models.AcademicPlan

	if err := r.db.Joins("Subject").Preload("AcademicYear").Find(&result, "subject.department_id = ? AND academic_year_id = ?", user.Department.ID, yearID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *AcademicPlanRepository) FindUserAcademicPlan(userUuid, uuid string) (*models.AcademicPlan, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var model models.AcademicPlan
	if err := r.db.Preload("AcademicYear").Joins("Subject").First(&model, "academic_plans.uuid = ? AND subject.department_id = ?", uuid, user.Department.ID).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *AcademicPlanRepository) FindSubjectBy(column string, value interface{}) (*models.Subject, error) {
	var model models.Subject
	if err := r.db.Preload("Department.User").First(&model, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *AcademicPlanRepository) FindAcademicYearBy(column string, value interface{}) (*models.AcademicYear, error) {
	var model models.AcademicYear
	if err := r.db.First(&model, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *AcademicPlanRepository) FindUser(userUuid string) (*models.User, error) {
	var result models.User
	if err := r.db.Preload("Department").First(&result, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *AcademicPlanRepository) FindBy(departmentID, yearID uint, column string, value interface{}) (*[]models.AcademicPlan, error) {
	var model []models.AcademicPlan
	if err := r.db.Preload("AcademicYear").Joins("Subject").Find(&model, fmt.Sprintf("academic_plans.%s = ? AND subject.department_id = ? AND academic_plans.academic_year_id = ?", column), value, departmentID, yearID).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *AcademicPlanRepository) CreateAcademicPlan(model *models.AcademicPlan) error {
	return r.db.Create(model).Error
}

func (r *AcademicPlanRepository) UpdateAcademicPlan(model *models.AcademicPlan) error {
	return r.db.Updates(model).Error
}

func (r *AcademicPlanRepository) UpdateOne(id uint, column string, value interface{}) error {
	return r.db.Model(&models.AcademicPlan{ID: id}).Update(column, value).Error
}

func (r *AcademicPlanRepository) DeleteAcademicPlan(model *models.AcademicPlan) error {
	return r.db.Delete(model).Error
}

func (r *AcademicPlanRepository) FindDepartmentBy(column string, value interface{}) (*models.Department, error) {
	var department models.Department
	if err := r.db.Preload("Major").Preload("Subjects").Preload("User.Role").First(&department, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &department, nil
}