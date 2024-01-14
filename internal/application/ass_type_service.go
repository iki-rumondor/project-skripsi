package application

import (
	"errors"
	"fmt"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"gorm.io/gorm"
)

type AssTypeService struct {
	Repo *repository.Repositories
}

func NewAssTypeService(repo *repository.Repositories) *AssTypeService {
	return &AssTypeService{
		Repo: repo,
	}
}

func (s *AssTypeService) CreateAssType(model *domain.AssessmentType) error {

	if err := s.Repo.AssTypeRepo.CreateAssType(model); err != nil {
		return &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}
	return nil
}

func (s *AssTypeService) GetAllAssType() (*[]domain.AssessmentType, error) {

	res, err := s.Repo.AssTypeRepo.FindAllAssType()
	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}

	return res, nil
}

func (s *AssTypeService) GetAssType(uuid string) (*domain.AssessmentType, error) {

	res, err := s.Repo.AssTypeRepo.FindAssTypeByUuid(uuid)
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

func (s *AssTypeService) UpdateAssType(AssType *domain.AssessmentType) error {

	if err := s.Repo.AssTypeRepo.UpdateAssType(AssType); err != nil {
		return &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}

	return nil
}

func (s *AssTypeService) DeleteAssType(AssType *domain.AssessmentType) error {

	if err := s.Repo.AssTypeRepo.DeleteAssType(AssType); err != nil {

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
