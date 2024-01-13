package application

import (
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
)

type UtilService struct {
	Repo *repository.Repositories
}

func NewUtilService(repo *repository.Repositories) *UtilService {
	return &UtilService{
		Repo: repo,
	}
}

func (s *UtilService) GetAllJurusan() (*[]domain.Jurusan, error) {

	jurusan, err := s.Repo.UtilRepo.FindAllJurusan()
	if err != nil {
		return nil, err
	}

	return jurusan, nil
}
