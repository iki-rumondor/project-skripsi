package repository

import (
	"fmt"

	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
)

type SubjectRepository struct {
	db *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) interfaces.SubjectRepoInterface {
	return &SubjectRepository{
		db: db,
	}
}

func (r *SubjectRepository) FindSubjects(userUuid string) (*[]models.Subject, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var subjects []models.Subject
	if err := r.db.Preload("Department").Find(&subjects, "department_id = ?", user.Department.ID).Error; err != nil {
		return nil, err
	}

	return &subjects, nil
}

func (r *SubjectRepository) FindUserSubject(userUuid, uuid string) (*models.Subject, error) {
	var user models.User
	if err := r.db.Preload("Department").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}
	
	var subject models.Subject
	if err := r.db.Preload("Department").First(&subject, "uuid = ? AND department_id = ?", uuid, user.Department.ID).Error; err != nil {
		return nil, err
	}

	return &subject, nil
}

func (r *SubjectRepository) FindUserBy(column string, value interface{}) (*models.User, error) {
	var result models.User
	if err := r.db.Preload("Department").First(&result, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *SubjectRepository) CreateSubject(subject *models.Subject) error {
	return r.db.Create(subject).Error
}

func (r *SubjectRepository) UpdateSubject(subject *models.Subject) error {
	if err := r.db.First(&models.Subject{}, "id = ? AND department_id = ?", subject.ID, subject.DepartmentID).Error; err != nil{
		return err
	}

	return r.db.Updates(subject).Error
}

func (r *SubjectRepository) DeleteSubject(subject *models.Subject) error {
	if err := r.db.First(&models.Subject{}, "id = ? AND department_id = ?", subject.ID, subject.DepartmentID).Error; err != nil{
		return err
	}

	return r.db.Delete(subject).Error
}
