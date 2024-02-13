package services

import (
	"errors"
	"log"

	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/http/response"
	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
)

type FacilityService struct {
	Repo interfaces.FacilityRepoInterface
}

func NewFacilityService(repo interfaces.FacilityRepoInterface) interfaces.FacilityServiceInterface {
	return &FacilityService{
		Repo: repo,
	}
}

func (s *FacilityService) CreateFacility(userUuid string, req *request.Facility) error {

	user, err := s.Repo.FindUserBy("uuid", userUuid)
	if err != nil {
		return err
	}

	model := models.Facility{
		DepartmentID: user.Department.ID,
		Name:         req.Name,
	}

	if err := s.Repo.CreateFacility(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}
	return nil
}

func (s *FacilityService) GetAllFacilities(userUuid string) (*[]models.Facility, error) {

	result, err := s.Repo.FindFacilities(userUuid)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *FacilityService) GetFacility(userUuid, uuid string) (*models.Facility, error) {
	result, err := s.Repo.FindUserFacility(userUuid, uuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("Facilitium Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *FacilityService) UpdateFacility(userUuid, uuid string, req *request.Facility) error {

	result, err := s.GetFacility(userUuid, uuid)
	if err != nil {
		return err
	}

	model := models.Facility{
		ID:           result.ID,
		DepartmentID: result.DepartmentID,
		Name:         req.Name,
	}

	if err := s.Repo.UpdateFacility(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *FacilityService) DeleteFacility(userUuid, uuid string) error {
	result, err := s.GetFacility(userUuid, uuid)
	if err != nil {
		return err
	}

	model := models.Facility{
		ID: result.ID,
	}

	if err := s.Repo.DeleteFacility(&model); err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return response.VIOLATED_ERR
		}
		return response.SERVICE_INTERR
	}

	return nil
}
