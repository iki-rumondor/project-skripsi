package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type AssTypeRepository interface {
	CreateAssType(model *domain.AssessmentType) error
	FindAllAssType() (*[]domain.AssessmentType, error)
	FindAssTypeByUuid(uuid string) (*domain.AssessmentType, error)
	UpdateAssType(model *domain.AssessmentType) error
	DeleteAssType(model *domain.AssessmentType) error
}
