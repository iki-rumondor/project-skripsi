package application

import (
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
)

type ProdiService struct {
	Repo *repository.Repositories
}

func NewProdiService(repo *repository.Repositories) *ProdiService {
	return &ProdiService{
		Repo: repo,
	}
}

func (s *ProdiService) GetAllProdi() (*[]domain.Prodi, error) {

	jurusan, err := s.Repo.ProdiRepo.FindAllProdi()
	if err != nil {
		return nil, err
	}

	return jurusan, nil
}

func (s *ProdiService) GetProdiByID(id uint) (*domain.Prodi, error) {

	jurusan, err := s.Repo.ProdiRepo.FindProdi(id)
	if err != nil {
		return nil, err
	}

	return jurusan, nil
}

func (s *ProdiService) CreateProdi(prodi *domain.Prodi) error {

	if err := s.Repo.ProdiRepo.CreateProdi(prodi); err != nil {
		return err
	}

	return nil
}

func (s *ProdiService) DeleteProdi(prodi *domain.Prodi) error {

	if err := s.Repo.ProdiRepo.DeleteProdi(prodi); err != nil {
		return err
	}

	return nil
}

func (s *ProdiService) UpdateProdi(prodi *domain.Prodi) error {

	if err := s.Repo.ProdiRepo.UpdateProdi(prodi); err != nil {
		return err
	}

	return nil
}
