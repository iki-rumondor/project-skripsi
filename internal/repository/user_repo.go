package repository

import (
	"fmt"

	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepoInterface {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindUserBy(column string, value interface{}) (*models.User, error) {
	var user models.User
	if err := r.db.Preload("Role").First(&user, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindSubjects() (*[]models.Subject, error) {
	var model []models.Subject
	if err := r.db.Find(&model).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *UserRepository) FindPracticalSubjects() (*[]models.Subject, error) {
	var model []models.Subject
	if err := r.db.Find(&model, "practical = 1").Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *UserRepository) CountMonevByYear(userUuid, yearUuid string) (map[string]int, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var year models.AcademicYear
	if err := r.db.First(&year, "uuid = ?", yearUuid).Error; err != nil {
		return nil, err
	}

	subjects := r.db.Model(&models.Subject{}).Where("department_id = ?", user.Department.ID).Select("id")
	fas := r.db.Model(&models.Facility{}).Where("department_id = ?", user.Department.ID).Select("id")
	teachers := r.db.Model(&models.Teacher{}).Where("department_id = ?", user.Department.ID).Select("id")

	var plans []models.AcademicPlan
	var modules []models.PracticalModule
	var tools []models.PracticalTool
	var skills []models.TeacherSkill
	var facilities []models.FacilityCondition

	r.db.Find(&plans, "subject_id IN (?) AND academic_year_id = ?", subjects, year.ID)
	r.db.Find(&modules, "subject_id IN (?) AND academic_year_id = ?", subjects, year.ID)
	r.db.Find(&tools, "subject_id IN (?) AND academic_year_id = ?", subjects, year.ID)
	r.db.Find(&skills, "teacher_id IN (?) AND academic_year_id = ?", teachers, year.ID)
	r.db.Find(&facilities, "facility_id IN (?) AND academic_year_id = ?", fas, year.ID)

	res := map[string]int{
		"plans": len(plans),
		"modules": len(modules),
		"tools": len(tools),
		"skills": len(skills),
		"facilities": len(facilities),
	}

	return res, nil
}