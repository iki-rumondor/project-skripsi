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

func (r *UserRepository) FindUsers() (*[]models.User, error) {
	var model []models.User
	if err := r.db.Preload(clause.Associations).Find(&model).Error; err != nil {
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

	var lastTeacherAttendeces []models.TeacherAttendence
	var lastStudentAttendeces []models.StudentAttendence
	var lastAcademicPlans []models.AcademicPlan

	var studentPassed []models.StudentAttendence
	var studentFinal []models.StudentAttendence
	var grade []models.TeacherAttendence

	r.db.Find(&plans, "subject_id IN (?) AND academic_year_id = ?", subjects, yearID)
	r.db.Find(&modules, "subject_id IN (?) AND academic_year_id = ?", subjects, yearID)
	r.db.Find(&tools, "subject_id IN (?) AND academic_year_id = ?", subjects, yearID)
	r.db.Find(&skills, "teacher_id IN (?) AND academic_year_id = ?", teachers, yearID)
	r.db.Find(&facilities, "facility_id IN (?) AND academic_year_id = ?", fas, yearID)

	r.db.Find(&teacherAttendeces, "subject_id IN (?) AND academic_year_id = ?", subjects, yearID)
	r.db.Find(&studentAttendeces, "subject_id IN (?) AND academic_year_id = ?", subjects, yearID)
	r.db.Find(&academicPlans, "subject_id IN (?) AND academic_year_id = ? AND middle = ?", subjects, yearID, true)

	r.db.Find(&lastTeacherAttendeces, "subject_id IN (?) AND academic_year_id = ? AND last != ?", subjects, yearID, 0)
	r.db.Find(&lastStudentAttendeces, "subject_id IN (?) AND academic_year_id = ? AND last != ?", subjects, yearID, 0)
	r.db.Find(&lastAcademicPlans, "subject_id IN (?) AND academic_year_id = ? AND last = ?", subjects, yearID, true)

	r.db.Find(&studentFinal, "subject_id IN (?) AND academic_year_id = ? AND final_amount != 0", subjects, yearID, 0)
	r.db.Find(&studentPassed, "subject_id IN (?) AND academic_year_id = ? AND passed_amount != 0", subjects, yearID, 0)
	r.db.Find(&grade, "subject_id IN (?) AND academic_year_id = ? AND grade_on_time = ?", subjects, yearID, true)

	res := map[string]int{
		"plans":      len(plans),
		"modules":    len(modules),
		"tools":      len(tools),
		"skills":     len(skills),
		"facilities": len(facilities),
		"t_att":      len(teacherAttendeces),
		"s_att":      len(studentAttendeces),
		"mid_plans":  len(academicPlans),
		"lt_att":     len(teacherAttendeces),
		"ls_att":     len(studentAttendeces),
		"last_plans": len(academicPlans),
		"passed":     len(studentPassed),
		"final":      len(studentFinal),
		"grade":      len(grade),
	}

	return res, nil
}

func (r *UserRepository) Update(id uint, tableName, column string, value interface{}) error {
	return r.db.Table(tableName).Where("id = ?", id).Update(column, value).Error
}

func (r *UserRepository) GetAll(tableName string) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	if err := r.db.Table(tableName).Find(&result).Error; err != nil {
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

func (r *UserRepository) CreateUser(model *models.User) error {
	return r.db.Create(model).Error
}

func (r *UserRepository) First(dest interface{}, condition string) error {
	return r.db.First(dest, condition).Error
}

func (r *UserRepository) FindDepartments(dest *[]models.Department) error {
	return r.db.Preload("Subjects.AcademicPlan").Find(dest).Error
}
