package repository

import (
	"fmt"

	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PracticalToolRepository struct {
	db *gorm.DB
}

func NewPracticalToolRepository(db *gorm.DB) interfaces.PracticalToolRepoInterface {
	return &PracticalToolRepository{
		db: db,
	}
}

func (r *PracticalToolRepository) FindPracticalTools(userUuid string, yearID uint) (*[]models.PracticalTool, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var result []models.PracticalTool

	if err := r.db.Joins("Subject").Preload("AcademicYear").Find(&result, "Subject.department_id = ? AND academic_year_id = ?", user.Department.ID, yearID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *PracticalToolRepository) FindUserPracticalTool(userUuid, uuid string) (*models.PracticalTool, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var model models.PracticalTool
	if err := r.db.Joins("Subject").Preload("AcademicYear").First(&model, "Practical_tools.uuid = ? AND Subject.department_id = ?", uuid, user.Department.ID).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *PracticalToolRepository) FindSubjectBy(column string, value interface{}) (*models.Subject, error) {
	var model models.Subject
	if err := r.db.Preload("Department.User").First(&model, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *PracticalToolRepository) FindAcademicYearBy(column string, value interface{}) (*models.AcademicYear, error) {
	var model models.AcademicYear
	if err := r.db.First(&model, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *PracticalToolRepository) CreatePracticalTool(model *models.PracticalTool) error {
	return r.db.Create(model).Error
}

func (r *PracticalToolRepository) UpdatePracticalTool(model *models.PracticalTool) error {
	return r.db.Updates(model).Error
}

func (r *PracticalToolRepository) DeletePracticalTool(model *models.PracticalTool) error {
	return r.db.Delete(model).Error
}

func (r *PracticalToolRepository) FindDepartment(uuid string) (*models.Department, error) {
	var department models.Department
	if err := r.db.First(&department, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &department, nil
}

func (r *PracticalToolRepository) FindByDepartment(departmentID, yearID uint) (*[]models.PracticalTool, error) {
	var result []models.PracticalTool

	if err := r.db.Joins("Subject").Preload(clause.Associations).Find(&result, "Subject.department_id = ? AND academic_year_id = ?", departmentID, yearID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
