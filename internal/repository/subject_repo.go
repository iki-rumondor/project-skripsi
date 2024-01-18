package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type SubjectRepository interface {
	CreateSubject(model *domain.Subject) error
	FindSubjectsWhere(column string, value interface{}) (*[]domain.Subject, error)
	FindAllSubject() (*[]domain.Subject, error)
	FindSubjectByUuid(uuid string) (*domain.Subject, error)
	UpdateSubject(model *domain.Subject) error
	DeleteSubject(model *domain.Subject) error
}
