package repository

import (
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type AssQuestionRepoImplementation struct {
	db *gorm.DB
}

func NewAssQuestionRepository(db *gorm.DB) AssQuestionRepository {
	return &AssQuestionRepoImplementation{
		db: db,
	}
}

func (r *AssQuestionRepoImplementation) CreateAssQuestion(model *domain.AssessmentQuestion) error {
	return r.db.Create(&model).Error
}

func (r *AssQuestionRepoImplementation) PaginationAssQuestion(pagination *domain.Pagination) (*domain.Pagination, error) {
	var results []domain.AssessmentQuestion
	var totalRows int64 = 0

	if err := r.db.Model(&domain.AssessmentQuestion{}).Count(&totalRows).Error; err != nil {
		return nil, err
	}

	if pagination.Limit == 0 {
		pagination.Limit = int(totalRows)
	}

	offset := pagination.Page * pagination.Limit

	if err := r.db.Limit(pagination.Limit).Offset(offset).Preload("AssessmentType").Find(&results).Error; err != nil {
		return nil, err
	}

	var res = []response.AssQuestion{}
	for _, item := range results {
		res = append(res, response.AssQuestion{
			Uuid:     item.Uuid,
			Question: item.Question,
			AssessmentType: &response.AssType{
				Uuid:      item.AssessmentType.Uuid,
				Type:      item.AssessmentType.Type,
				CreatedAt: item.AssessmentType.CreatedAt,
				UpdatedAt: item.AssessmentType.UpdatedAt,
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	pagination.Rows = res

	pagination.TotalRows = int(totalRows)

	return pagination, nil
}

func (r *AssQuestionRepoImplementation) FindAllAssQuestion() (*[]domain.AssessmentQuestion, error) {
	var result []domain.AssessmentQuestion
	if err := r.db.Preload("AssessmentType").Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *AssQuestionRepoImplementation) FindAssQuestionByUuid(uuid string) (*domain.AssessmentQuestion, error) {
	var result domain.AssessmentQuestion
	if err := r.db.Preload("AssessmentType").First(&result, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *AssQuestionRepoImplementation) UpdateAssQuestion(model *domain.AssessmentQuestion) error {
	return r.db.Model(&model).Updates(&model).Error
}

func (r *AssQuestionRepoImplementation) DeleteAssQuestion(model *domain.AssessmentQuestion) error {
	return r.db.Delete(&model).Error
}
