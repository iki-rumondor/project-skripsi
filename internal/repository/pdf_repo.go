package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/iki-rumondor/go-monev/internal/http/response"
	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PdfRepository struct {
	db *gorm.DB
}

func NewPdfRepository(db *gorm.DB) interfaces.PdfRepoInterface {
	return &PdfRepository{
		db: db,
	}
}

func (r *PdfRepository) LaravelPdf(data interface{}, endpoint string) (*http.Response, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var API_URL = os.Getenv("LARAVEL_PDF")
	if API_URL == "" {
		return nil, response.SERVICE_INTERR
	}

	url := fmt.Sprintf("%s/%s", API_URL, endpoint)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *PdfRepository) FindDepartmentBy(column string, value interface{}) (*models.Department, error) {
	var department models.Department
	if err := r.db.Preload("Major").Preload("Subjects").Preload("User.Role").First(&department, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &department, nil
}

func (r *PdfRepository) FindAcademicYearBy(column string, value interface{}) (*models.AcademicYear, error) {
	var model models.AcademicYear
	if err := r.db.First(&model, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *PdfRepository) FindAcademicPlans(departmentID, yearID uint) (*[]models.AcademicPlan, error) {

	var result []models.AcademicPlan

	if err := r.db.Joins("Subject").Preload("AcademicYear").Find(&result, "Subject.department_id = ? AND academic_year_id = ?", departmentID, yearID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *PdfRepository) FindModules(departmentID, yearID uint) (*[]models.PracticalModule, error) {
	var result []models.PracticalModule

	if err := r.db.Joins("Subject").Preload("AcademicYear").Preload("Laboratory").Find(&result, "Subject.department_id = ? AND academic_year_id = ?", departmentID, yearID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *PdfRepository) FindTools(departmentID, yearID uint) (*[]models.PracticalTool, error) {
	var result []models.PracticalTool

	if err := r.db.Joins("Subject").Preload(clause.Associations).Find(&result, "Subject.department_id = ? AND academic_year_id = ?", departmentID, yearID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *PdfRepository) FindSkills(departmentID, yearID uint) (*[]models.TeacherSkill, error) {
	var result []models.TeacherSkill

	if err := r.db.Joins("Subject").Preload(clause.Associations).Find(&result, "Subject.department_id = ? AND academic_year_id = ?", departmentID, yearID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *PdfRepository) FindFacilities(departmentID, yearID uint) (*[]models.FacilityCondition, error) {
	var result []models.FacilityCondition

	if err := r.db.Joins("Facility").Preload(clause.Associations).Find(&result, "Facility.department_id = ? AND academic_year_id = ?", departmentID, yearID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *PdfRepository) FindStudentAttendences(departmentID, yearID uint) (*[]models.StudentAttendence, error) {
	var result []models.StudentAttendence
	if err := r.db.Joins("Subject").Preload("AcademicYear").Find(&result, "Subject.department_id = ? AND academic_year_id = ?", departmentID, yearID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *PdfRepository) FindTeacherAttendences(departmentID, yearID uint) (*[]models.TeacherAttendence, error) {
	var result []models.TeacherAttendence
	if err := r.db.Joins("Subject").Preload(clause.Associations).Find(&result, "Subject.department_id = ? AND academic_year_id = ?", departmentID, yearID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
