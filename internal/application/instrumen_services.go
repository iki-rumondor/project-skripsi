package application

import (
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
)

type InstrumenService struct {
	Repo *repository.Repositories
}

func NewInstrumenService(repo *repository.Repositories) *InstrumenService {
	return &InstrumenService{
		Repo: repo,
	}
}

func (s *InstrumenService) GetAllIndikator() (*[]domain.Indikator, error) {

	res, err := s.Repo.InstrumenRepo.FindAllIndikator()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *InstrumenService) GetIndikator(id uint) (*domain.Indikator, error) {

	data := domain.Indikator{
		ID: id,
	}

	res, err := s.Repo.InstrumenRepo.FindIndikator(&data)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *InstrumenService) CreateIndikator(req *request.Indikator) error {
	model := domain.Indikator{
		Description:     req.Description,
		IndikatorTypeID: req.TypeID,
	}

	if err := s.Repo.InstrumenRepo.CreateIndikator(&model); err != nil {
		return err
	}

	return nil
}

func (s *InstrumenService) UpdateIndikator(id uint, req *request.Indikator) error {
	model := domain.Indikator{
		ID:              id,
		Description:     req.Description,
		IndikatorTypeID: req.TypeID,
	}

	if err := s.Repo.InstrumenRepo.UpdateIndikator(&model); err != nil {
		return err
	}

	return nil
}

func (s *InstrumenService) DeleteIndikator(id uint) error {

	data := &domain.Indikator{
		ID: id,
	}

	if err := s.Repo.InstrumenRepo.DeleteIndikator(data); err != nil {
		return err
	}

	return nil
}

func (s *InstrumenService) CreateInstrumenType(req *request.CreateInstrumenType) error {
	model := domain.InstrumenType{
		Name: req.Name,
	}

	if err := s.Repo.InstrumenRepo.CreateInstrumenType(&model); err != nil {
		return err
	}

	return nil
}

func (s *InstrumenService) GetAllInstrumenType() (*[]domain.InstrumenType, error) {

	res, err := s.Repo.InstrumenRepo.FindAllInstrumenType()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *InstrumenService) DeleteInstrumenType(id uint) error {

	data := &domain.InstrumenType{
		ID: id,
	}

	if err := s.Repo.InstrumenRepo.DeleteInstrumenType(data); err != nil {
		return err
	}

	return nil
}

func (s *InstrumenService) GetAllIndikatorType() (*[]domain.IndikatorType, error) {

	res, err := s.Repo.InstrumenRepo.FindAllIndikatorType()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *InstrumenService) CreateIndikatorType(req *request.CreateIndikatorType) error {
	model := domain.IndikatorType{
		Description: req.Description,
		TypeID:      req.InstrumentID,
	}

	if err := s.Repo.InstrumenRepo.CreateIndikatorType(&model); err != nil {
		return err
	}

	return nil
}

func (s *InstrumenService) DeleteIndikatorType(id uint) error {

	data := &domain.IndikatorType{
		ID: id,
	}

	if err := s.Repo.InstrumenRepo.DeleteIndikatorType(data); err != nil {
		return err
	}

	return nil
}
