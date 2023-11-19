package application

import (
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
)

type UtilService struct {
	Repo repository.UtilRepository
}

func NewUtilService(repo repository.UtilRepository) *UtilService {
	return &UtilService{
		Repo: repo,
	}
}

func (s *UtilService) GetAllJurusan() (*[]domain.Jurusan, error) {

	jurusan, err := s.Repo.FindAllJurusan()
	if err != nil {
		return nil, err
	}

	return jurusan, nil
}
