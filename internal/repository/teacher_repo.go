package repository

import (
	"fmt"

	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
)

type TeacherRepository struct {
	db *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) interfaces.TeacherRepoInterface {
	return &TeacherRepository{
		db: db,
	}
}

func (r *TeacherRepository) FindTeachers(userUuid string) (*[]models.Teacher, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var Teachers []models.Teacher
	if err := r.db.Preload("Department").Find(&Teachers, "department_id = ?", user.Department.ID).Error; err != nil {
		return nil, err
	}

	return &Teachers, nil
}

func (r *TeacherRepository) FindUserTeacher(userUuid, uuid string) (*models.Teacher, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var Teacher models.Teacher
	if err := r.db.Preload("Department").First(&Teacher, "uuid = ? AND department_id = ?", uuid, user.Department.ID).Error; err != nil {
		return nil, err
	}

	return &Teacher, nil
}

func (r *TeacherRepository) FindUserBy(column string, value interface{}) (*models.User, error) {
	var result models.User
	if err := r.db.Preload("Department").First(&result, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *TeacherRepository) CreateTeacher(Teacher *models.Teacher) error {
	return r.db.Create(Teacher).Error
}

func (r *TeacherRepository) UpdateTeacher(Teacher *models.Teacher) error {
	if err := r.db.First(&models.Teacher{}, "id = ? AND department_id = ?", Teacher.ID, Teacher.DepartmentID).Error; err != nil {
		return err
	}

	return r.db.Updates(Teacher).Error
}

func (r *TeacherRepository) DeleteTeacher(Teacher *models.Teacher) error {
	if err := r.db.First(&models.Teacher{}, "id = ? AND department_id = ?", Teacher.ID, Teacher.DepartmentID).Error; err != nil {
		return err
	}

	return r.db.Delete(Teacher).Error
}
