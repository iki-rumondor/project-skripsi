package application

import (
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
)

type AdminService struct {
	InstrumenRepo repository.InstrumenRepository
}

func NewAdminService(repositories repository.Repositories) *AdminService {
	return &AdminService{
		InstrumenRepo: repositories.InstrumenRepository,
	}
}

func (s *AdminService) GetAllIndikator() (*[]domain.Indikator, error) {

	res, err := s.InstrumenRepo.FindAllIndikator()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *AdminService) GetIndikator(id uint) (*domain.Indikator, error) {

	data := domain.Indikator{
		ID: id,
	}

	res, err := s.InstrumenRepo.FindIndikator(&data)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *AdminService) CreateIndikator(req *request.Indikator) error {
	model := domain.Indikator{
		Description: req.Description,
		IndikatorTypeID: req.TypeID,
	}

	if err := s.InstrumenRepo.CreateIndikator(&model); err != nil {
		return err
	}

	return nil
}

func (s *AdminService) UpdateIndikator(id uint, req *request.Indikator) error {
	model := domain.Indikator{
		ID: id,
		Description: req.Description,
		IndikatorTypeID: req.TypeID,
	}

	if err := s.InstrumenRepo.UpdateIndikator(&model); err != nil {
		return err
	}

	return nil
}

func (s *AdminService) DeleteIndikator(id uint) (error) {

	data := &domain.Indikator{
		ID: id,
	}

	if err := s.InstrumenRepo.DeleteIndikator(data); err != nil {
		return err
	}

	return nil
}

func (s *AdminService) CreateInstrumenType(req *request.CreateInstrumenType) error {
	model := domain.InstrumenType{
		Name: req.Name,
	}

	if err := s.InstrumenRepo.CreateInstrumenType(&model); err != nil {
		return err
	}

	return nil
}

func (s *AdminService) GetAllInstrumenType() (*[]domain.InstrumenType, error) {

	res, err := s.InstrumenRepo.FindAllInstrumenType()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *AdminService) DeleteInstrumenType(id uint) (error) {

	data := &domain.InstrumenType{
		ID: id,
	}

	if err := s.InstrumenRepo.DeleteInstrumenType(data); err != nil {
		return err
	}

	return nil
}

func (s *AdminService) GetAllIndikatorType() (*[]domain.IndikatorType, error) {

	res, err := s.InstrumenRepo.FindAllIndikatorType()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *AdminService) CreateIndikatorType(req *request.CreateIndikatorType) error {
	model := domain.IndikatorType{
		Description: req.Description,
		TypeID: req.InstrumentID,
	}

	if err := s.InstrumenRepo.CreateIndikatorType(&model); err != nil {
		return err
	}

	return nil
}

func (s *AdminService) DeleteIndikatorType(id uint) (error) {

	data := &domain.IndikatorType{
		ID: id,
	}

	if err := s.InstrumenRepo.DeleteIndikatorType(data); err != nil {
		return err
	}

	return nil
}