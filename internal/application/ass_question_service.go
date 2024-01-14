package application

import (
	"errors"
	"fmt"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"gorm.io/gorm"
)

type AssQuestionService struct {
	Repo *repository.Repositories
}

func NewAssQuestionService(repo *repository.Repositories) *AssQuestionService {
	return &AssQuestionService{
		Repo: repo,
	}
}

func (s *AssQuestionService) CreateAssQuestion(model *domain.AssessmentQuestion) error {

	if err := s.Repo.AssQuestionRepo.CreateAssQuestion(model); err != nil {
		return &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}
	return nil
}

func (s *AssQuestionService) GetAllAssQuestion(urlPath string, pagination *domain.Pagination) (*domain.Pagination, error) {

	res, err := s.Repo.AssQuestionRepo.PaginationAssQuestion(pagination)
	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}

	page := GeneratePages(urlPath, res)

	return page, nil
}

func (s *AssQuestionService) GetAssQuestion(uuid string) (*domain.AssessmentQuestion, error) {

	res, err := s.Repo.AssQuestionRepo.FindAssQuestionByUuid(uuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.Error{
				Code:    404,
				Message: fmt.Sprintf("Tipe penilaian dengan uuid %s tidak ditemukan", uuid),
			}
		}
		return nil, &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}

	return res, nil
}

func (s *AssQuestionService) UpdateAssQuestion(AssQuestion *domain.AssessmentQuestion) error {

	if err := s.Repo.AssQuestionRepo.UpdateAssQuestion(AssQuestion); err != nil {
		return &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}

	return nil
}

func (s *AssQuestionService) DeleteAssQuestion(AssQuestion *domain.AssessmentQuestion) error {

	if err := s.Repo.AssQuestionRepo.DeleteAssQuestion(AssQuestion); err != nil {

		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return &response.Error{
				Code:    500,
				Message: "Data ini tidak dapat dihapus karena berelasi dengan data lain",
			}
		}
		return &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}

	return nil
}
