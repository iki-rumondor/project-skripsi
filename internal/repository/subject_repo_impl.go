package repository

import (
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type SubjectRepoImplementation struct {
	db *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) SubjectRepository {
	return &SubjectRepoImplementation{
		db: db,
	}
}

func (r *SubjectRepoImplementation) CreateSubject(model *domain.Subject) error {
	return r.db.Create(&model).Error
}

func (r *SubjectRepoImplementation) FindAllSubject() (*[]domain.Subject, error) {
	var result []domain.Subject
	if err := r.db.Preload("Prodi").Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *SubjectRepoImplementation) FindSubjectByUuid(uuid string) (*domain.Subject, error) {
	var result domain.Subject
	if err := r.db.Preload("Prodi").First(&result, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *SubjectRepoImplementation) UpdateSubject(model *domain.Subject) error {
	return r.db.Model(&model).Updates(&model).Error
}

func (r *SubjectRepoImplementation) DeleteSubject(model *domain.Subject) error {
	return r.db.Delete(&model).Error
}
