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

func (r *SubjectRepository) FindSubjects() (*[]models.Subject, error) {
	var subjects []models.Subject
	if err := r.db.Preload("Department").Find(&subjects).Error; err != nil {
		return nil, err
	}

	return &subjects, nil
}

func (r *SubjectRepository) FindSubjectBy(column string, value interface{}) (*models.Subject, error) {
	var subject models.Subject
	if err := r.db.Preload("Department").First(&subject, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &subject, nil
}

func (r *SubjectRepository) FindDepartmentBy(column string, value interface{}) (*models.Department, error) {
	var department models.Department
	if err := r.db.First(&department, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &department, nil
}

func (r *SubjectRepository) UpsertSubject(subject *models.Subject) error {
	return r.db.Save(&subject).Error
}

func (r *SubjectRepository) DeleteSubject(subject *models.Subject) error {
	return r.db.Delete(subject).Error
}
