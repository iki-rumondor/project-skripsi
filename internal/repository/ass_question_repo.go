package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type AssQuestionRepository interface {
	CreateAssQuestion(model *domain.AssessmentQuestion) error
	PaginationAssQuestion(pagination *domain.Pagination) (*domain.Pagination, error)
	FindAllAssQuestion() (*[]domain.AssessmentQuestion, error)
	FindAssQuestionByUuid(uuid string) (*domain.AssessmentQuestion, error)
	UpdateAssQuestion(model *domain.AssessmentQuestion) error
	DeleteAssQuestion(model *domain.AssessmentQuestion) error
}
