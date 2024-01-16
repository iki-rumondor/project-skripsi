package repository

import (
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type ResponseRepoImplementation struct {
	db *gorm.DB
}

func NewResponseRepository(db *gorm.DB) ResponseRepository {
	return &ResponseRepoImplementation{
		db: db,
	}
}

func (r *ResponseRepoImplementation) CreateResponse(model *domain.Response) error {
	return r.db.Create(&model).Error
}

func (r *ResponseRepoImplementation) PaginationResponse(pagination *domain.Pagination) (*domain.Pagination, error) {
	var results []domain.Response
	var totalRows int64 = 0

	if err := r.db.Model(&domain.Response{}).Count(&totalRows).Error; err != nil {
		return nil, err
	}

	if pagination.Limit == 0 {
		pagination.Limit = int(totalRows)
	}

	offset := pagination.Page * pagination.Limit

	if err := r.db.Limit(pagination.Limit).Offset(offset).Preload("User").Preload("Question").Find(&results).Error; err != nil {
		return nil, err
	}

	var res = []response.AssResponse{}
	for _, item := range results {
		res = append(res, response.AssResponse{
			Uuid:     item.Uuid,
			Response: item.Response,
			User: &response.User{
				Uuid:     item.User.Uuid,
				Username: item.User.Username,
				Role:     item.User.Role,
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	pagination.Rows = res

	pagination.TotalRows = int(totalRows)

	return pagination, nil
}

func (r *ResponseRepoImplementation) FindAllResponse() (*[]domain.Response, error) {
	var result []domain.Response
	if err := r.db.Preload("User").Preload("Question").Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *ResponseRepoImplementation) FindResponseByUuid(uuid string) (*domain.Response, error) {
	var result domain.Response
	if err := r.db.Preload("User").Preload("Question").First(&result, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *ResponseRepoImplementation) UpdateResponse(model *domain.Response) error {
	return r.db.Model(&model).Updates(&model).Error
}

func (r *ResponseRepoImplementation) DeleteResponse(model *domain.Response) error {
	return r.db.Delete(&model).Error
}
