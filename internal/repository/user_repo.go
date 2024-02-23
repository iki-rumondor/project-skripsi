package repository

import (
	"fmt"

	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	if err := r.db.Preload(clause.Associations).First(&user, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
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

func (r *UserRepository) CountMonevByYear(departmentID, yearID uint) (map[string]int, error) {

	subjects := r.db.Model(&models.Subject{}).Where("department_id = ?", departmentID).Select("id")
	fas := r.db.Model(&models.Facility{}).Where("department_id = ?", departmentID).Select("id")
	teachers := r.db.Model(&models.Teacher{}).Where("department_id = ?", departmentID).Select("id")

	var plans []models.AcademicPlan
	var modules []models.PracticalModule
	var tools []models.PracticalTool
	var skills []models.TeacherSkill
	var facilities []models.FacilityCondition

	var teacherAttendeces []models.TeacherAttendence
	var studentAttendeces []models.StudentAttendence
	var academicPlans []models.AcademicPlan

	r.db.Find(&plans, "subject_id IN (?) AND academic_year_id = ?", subjects, yearID)
	r.db.Find(&modules, "subject_id IN (?) AND academic_year_id = ?", subjects, yearID)
	r.db.Find(&tools, "subject_id IN (?) AND academic_year_id = ?", subjects, yearID)
	r.db.Find(&skills, "teacher_id IN (?) AND academic_year_id = ?", teachers, yearID)
	r.db.Find(&facilities, "facility_id IN (?) AND academic_year_id = ?", fas, yearID)

	r.db.Find(&teacherAttendeces, "subject_id IN (?) AND academic_year_id = ?", subjects, yearID)
	r.db.Find(&studentAttendeces, "subject_id IN (?) AND academic_year_id = ?", subjects, yearID)
	r.db.Find(&academicPlans, "subject_id IN (?) AND academic_year_id = ? AND available = ?", subjects, yearID, true)

	res := map[string]int{
		"plans":      len(plans),
		"modules":    len(modules),
		"tools":      len(tools),
		"skills":     len(skills),
		"facilities": len(facilities),
		"t_att":      len(teacherAttendeces),
		"s_att":      len(studentAttendeces),
		"av_plans":   len(academicPlans),
	}

	return res, nil
}

func (r *UserRepository) Update(id uint, tableName, column string, value interface{}) error {
	return r.db.Table(tableName).Where("id = ?", id).Update(column, value).Error
}

func (r *UserRepository) GetAll(tableName string) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	if err := r.db.Table(tableName).Take(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *UserRepository) GetOne(tableName, column string, value interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}
	if err := r.db.Table(tableName).Where(fmt.Sprintf("%s = ?", column), value).Take(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
