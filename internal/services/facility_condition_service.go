package services

import (
	"errors"
	"log"
	"strconv"

	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/http/response"
	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
)

type FacilityConditionService struct {
	Repo interfaces.FacilityConditionRepoInterface
}

func NewFacilityConditionService(repo interfaces.FacilityConditionRepoInterface) interfaces.FacilityConditionServiceInterface {
	return &FacilityConditionService{
		Repo: repo,
	}
}

func (s *FacilityConditionService) CreateFacilityCondition(userUuid string, req *request.FacilityCondition) error {

	facility, err := s.Repo.FindFacilityBy("uuid", req.FacilityUuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Fasilitas Tidak Ditemukan")
		}
		return response.SERVICE_INTERR
	}

	if facility.Department.User.Uuid != userUuid {
		return response.NOTFOUND_ERR("Fasilitas Tidak Ditemukan")
	}

	academic_year, err := s.Repo.FindAcademicYearBy("uuid", req.AcademicYearUuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Tahun Ajaran Tidak Ditemukan")
		}
		return response.SERVICE_INTERR
	}

	amount, _ := strconv.Atoi(req.Amount)
	deactive, _ := strconv.Atoi(req.Deactive)

	if deactive > amount {
		return response.BADREQ_ERR("Jumlah Rusak Tidak Boleh Lebih Dari Jumlah Barang")
	}

	model := models.FacilityCondition{
		Amount:         uint(amount),
		Unit:           req.Unit,
		Deactive:       uint(deactive),
		Note:           req.Note,
		FacilityID:     facility.ID,
		AcademicYearID: academic_year.ID,
	}

	if err := s.Repo.CreateFacilityCondition(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}
	return nil
}

func (s *FacilityConditionService) GetFacilityConditionsByYear(userUuid, yearUuid string) (*[]models.FacilityCondition, error) {
	year, err := s.Repo.FindAcademicYearBy("uuid", yearUuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("Tahun Ajaran Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	result, err := s.Repo.FindFacilityConditionsByYear(userUuid, year.ID)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *FacilityConditionService) GetFacilityOptions(userUuid, yearUuid string) (*[]models.Facility, error) {
	year, err := s.Repo.FindAcademicYearBy("uuid", yearUuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("Tahun Ajaran Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	result, err := s.Repo.FindFacilityOptions(userUuid, year.ID)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *FacilityConditionService) GetFacilityCondition(userUuid, uuid string) (*models.FacilityCondition, error) {
	result, err := s.Repo.FindUserFacilityCondition(userUuid, uuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("RPS Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *FacilityConditionService) UpdateFacilityCondition(userUuid, uuid string, req *request.UpdateFacilityCondition) error {

	result, err := s.GetFacilityCondition(userUuid, uuid)
	if err != nil {
		return err
	}

	amount, _ := strconv.Atoi(req.Amount)
	deactive, _ := strconv.Atoi(req.Deactive)

	if deactive > amount {
		return response.BADREQ_ERR("Jumlah Rusak Tidak Boleh Lebih Dari Jumlah Barang")
	}

	model := models.FacilityCondition{
		ID:       result.ID,
		Amount:   uint(amount),
		Unit:     req.Unit,
		Deactive: uint(deactive),
		Note:     req.Note,
	}
	if err := s.Repo.UpdateFacilityCondition(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *FacilityConditionService) DeleteFacilityCondition(userUuid, uuid string) error {
	result, err := s.GetFacilityCondition(userUuid, uuid)
	if err != nil {
		return err
	}

	model := models.FacilityCondition{
		ID: result.ID,
	}

	if err := s.Repo.DeleteFacilityCondition(&model); err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return response.VIOLATED_ERR
		}
		return response.SERVICE_INTERR
	}

	return nil
}
