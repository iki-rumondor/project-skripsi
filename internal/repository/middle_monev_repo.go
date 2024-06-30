package repository

import (
	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MiddleMonevRepository struct {
	db *gorm.DB
}

func NewMiddleMonevRepository(db *gorm.DB) interfaces.MiddleMonevRepoInterface {
	return &MiddleMonevRepository{
		db: db,
	}
}

func (r *MiddleMonevRepository) CreateTeacherAttendence(model *models.TeacherAttendence) error {
	return r.db.Create(model).Error
}

func (r *MiddleMonevRepository) FindTeacherAttendences(userUuid string, yearID uint) (*[]models.TeacherAttendence, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var result []models.TeacherAttendence
	if err := r.db.Joins("Subject").Preload("Teacher").Preload("AcademicYear").Order("subject_id desc, class asc").Find(&result, "Subject.department_id = ? AND academic_year_id = ?", user.Department.ID, yearID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *MiddleMonevRepository) FindTeacherAttendence(userUuid, uuid string) (*models.TeacherAttendence, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var result models.TeacherAttendence
	if err := r.db.Joins("Subject").Preload("Teacher").Preload("AcademicYear").First(&result, "Subject.department_id = ? AND teacher_attendences.uuid = ?", user.Department.ID, uuid).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *MiddleMonevRepository) CreateStudentAttendence(model *models.StudentAttendence) error {
	return r.db.Create(model).Error
}

func (r *MiddleMonevRepository) FindStudentAttendences(userUuid string, yearID uint) (*[]models.StudentAttendence, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var result []models.StudentAttendence
	if err := r.db.Joins("Subject").Preload("AcademicYear").Order("subject_id asc, class asc").Find(&result, "Subject.department_id = ? AND academic_year_id = ?", user.Department.ID, yearID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *MiddleMonevRepository) FindStudentAttendence(userUuid, uuid string) (*models.StudentAttendence, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var result models.StudentAttendence
	if err := r.db.Joins("Subject").Preload("AcademicYear").First(&result, "Subject.department_id = ? AND Student_attendences.uuid = ?", user.Department.ID, uuid).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *MiddleMonevRepository) FindAcademicYear(uuid string) (*models.AcademicYear, error) {
	var result models.AcademicYear
	if err := r.db.First(&result, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *MiddleMonevRepository) FindTeacher(uuid string) (*models.Teacher, error) {
	var result models.Teacher
	if err := r.db.First(&result, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *MiddleMonevRepository) FindUserSubject(userUuid, uuid string) (*models.Subject, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var result models.Subject
	if err := r.db.Preload("Department").Preload("AcademicPlan.AcademicYear").First(&result, "uuid = ? AND department_id = ?", uuid, user.Department.ID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *MiddleMonevRepository) CountTotalMonev(userUuid string, yearID uint) (map[string]int, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	subjects := r.db.Model(&models.Subject{}).Where("department_id = ?", user.Department.ID).Select("id")

	var teacherAttendeces []models.TeacherAttendence
	var studentAttendeces []models.StudentAttendence
	var academicPlans []models.AcademicPlan
	var lastTeacherAttendeces []models.TeacherAttendence
	var lastStudentAttendeces []models.StudentAttendence

	r.db.Find(&teacherAttendeces, "subject_id IN (?) AND academic_year_id = ?", subjects, yearID)
	r.db.Find(&studentAttendeces, "subject_id IN (?) AND academic_year_id = ?", subjects, yearID)
	r.db.Find(&academicPlans, "subject_id IN (?) AND academic_year_id = ? AND available = ?", subjects, yearID, true)

	r.db.Find(&lastTeacherAttendeces, "subject_id IN (?) AND academic_year_id = ? AND middle = ?", subjects, yearID, true)
	r.db.Find(&lastStudentAttendeces, "subject_id IN (?) AND academic_year_id = ? AND middle = ?", subjects, yearID, true)

	res := map[string]int{
		"t_att":      len(teacherAttendeces),
		"s_att":      len(studentAttendeces),
		"plans":      len(academicPlans),
		"t_att_last": len(lastTeacherAttendeces),
		"s_att_last": len(lastStudentAttendeces),
	}

	return res, nil
}

func (r *MiddleMonevRepository) UpdateTeacherAttendence(model *models.TeacherAttendence) error {
	return r.db.Updates(model).Error
}

func (r *MiddleMonevRepository) UpdateStudentAttendence(model *models.StudentAttendence) error {
	return r.db.Updates(model).Error
}

func (r *MiddleMonevRepository) DeleteTeacherAttendence(model *models.TeacherAttendence) error {
	return r.db.Delete(model).Error
}

func (r *MiddleMonevRepository) DeleteStudentAttendence(model *models.StudentAttendence) error {
	return r.db.Delete(model).Error
}

func (r *MiddleMonevRepository) FindDepartment(uuid string) (*models.Department, error) {
	var department models.Department
	if err := r.db.Preload(clause.Associations).First(&department, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &department, nil
}
