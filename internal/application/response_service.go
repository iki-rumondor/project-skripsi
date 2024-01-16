package application

import (
	"errors"
	"fmt"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"gorm.io/gorm"
)

type ResponseService struct {
	Repo *repository.Repositories
}

func NewResponseService(repo *repository.Repositories) *ResponseService {
	return &ResponseService{
		Repo: repo,
	}
}

func (s *ResponseService) CreateResponse(model *domain.Response) error {

	if err := s.Repo.ResponseRepo.CreateResponse(model); err != nil {
		return &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}
	return nil
}

func (s *ResponseService) GetAllResponse(urlPath string, pagination *domain.Pagination) (*domain.Pagination, error) {

	res, err := s.Repo.ResponseRepo.PaginationResponse(pagination)
	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}

	page := GeneratePages(urlPath, res)

	return page, nil
}

func (s *ResponseService) GetResponse(uuid string) (*domain.Response, error) {

	res, err := s.Repo.ResponseRepo.FindResponseByUuid(uuid)
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

func (s *ResponseService) UpdateResponse(Response *domain.Response) error {

	if err := s.Repo.ResponseRepo.UpdateResponse(Response); err != nil {
		return &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}

	return nil
}

func (s *ResponseService) DeleteResponse(Response *domain.Response) error {

	if err := s.Repo.ResponseRepo.DeleteResponse(Response); err != nil {

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
