package repository

import (
	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
)

type LastMonevRepository struct {
	db *gorm.DB
}

func NewLastMonevRepository(db *gorm.DB) interfaces.LastMonevRepoInterface {
	return &LastMonevRepository{
		db: db,
	}
}

func (r *LastMonevRepository) CountLastMonev(departmentID, yearID uint) (map[string]int, error) {
	subjects := r.db.Model(&models.Subject{}).Where("department_id = ?", departmentID).Select("id")

	var studentAttendence []models.StudentAttendence
	var teachers []models.TeacherAttendence

	r.db.Find(&studentAttendence, "subject_id IN (?) AND academic_year_id = ?", subjects, yearID)
	r.db.Find(&teachers, "subject_id IN (?) AND academic_year_id = ?", subjects, yearID)

	res := map[string]int{
		"s_att": len(studentAttendence),
		"t_att": len(teachers),
	}

	return res, nil
}

func (r *LastMonevRepository) UpdateStudentAttendence(model *models.StudentAttendence, column string, value interface{}) error {
	return r.db.Model(model).Update(column, value).Error
}

func (r *LastMonevRepository) UpdateTeacherAttendence(model *models.TeacherAttendence, column string, value interface{}) error {
	return r.db.Model(model).Update(column, value).Error
}

func (r *LastMonevRepository) FindStudentAttendence(userUuid, uuid string) (*models.StudentAttendence, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var result models.StudentAttendence
	if err := r.db.Joins("Subject").Preload("AcademicYear").First(&result, "subject.department_id = ? AND student_attendences.uuid = ?", user.Department.ID, uuid).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *LastMonevRepository) FindTeacherAttendence(userUuid, uuid string) (*models.TeacherAttendence, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var result models.TeacherAttendence
	if err := r.db.Joins("Subject").Preload("AcademicYear").First(&result, "subject.department_id = ? AND teacher_attendences.uuid = ?", user.Department.ID, uuid).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *LastMonevRepository) FindAcademicYear(uuid string) (*models.AcademicYear, error) {
	var result models.AcademicYear
	if err := r.db.First(&result, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *LastMonevRepository) FindUser(uuid string) (*models.User, error) {
	var result models.User
	if err := r.db.Preload("Department").First(&result, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *LastMonevRepository) FindUserSubject(userUuid, uuid string) (*models.Subject, error) {
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
