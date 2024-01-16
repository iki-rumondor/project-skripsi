package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type ResponseRepository interface {
	CreateResponse(model *domain.Response) error
	PaginationResponse(pagination *domain.Pagination) (*domain.Pagination, error)
	FindAllResponse() (*[]domain.Response, error)
	FindResponseByUuid(uuid string) (*domain.Response, error)
	UpdateResponse(model *domain.Response) error
	DeleteResponse(model *domain.Response) error
}
