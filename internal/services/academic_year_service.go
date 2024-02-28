package services

import (
	"errors"
	"log"

	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/http/response"
	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"github.com/iki-rumondor/go-monev/internal/utils"
	"gorm.io/gorm"
)

type AcademicYearService struct {
	Repo interfaces.AcademicYearRepoInterface
}

func NewAcademicYearService(repo interfaces.AcademicYearRepoInterface) interfaces.AcademicYearServiceInterface {
	return &AcademicYearService{
		Repo: repo,
	}
}

func (s *AcademicYearService) CreateAcademicYear(req *request.AcademicYear) error {

	// first := utils.AddDate(req.FirstDate, req.FirstAmount)
	// middle := utils.AddDate(req.MiddleDate, req.MiddleAmount)
	// middle_last := utils.AddDate(req.MiddleLastDate, req.MiddleLastAmount)
	// last := utils.AddDate(req.LastDate, req.LastAmount)

	model := models.AcademicYear{
		Semester:       req.Semester,
		Year:           req.Year,
		FirstDate:      req.FirstDate,
		MiddleDate:     req.MiddleDate,
		MiddleLastDate: req.MiddleLastDate,
		LastDate:       req.LastDate,
		FirstDays:      req.FirstDays,
		MiddleDays:     req.MiddleDays,
		MiddleLastDays: req.MiddleLastDays,
		LastDays:       req.LastDays,
	}

	if err := s.Repo.CreateAcademicYear(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}
	return nil
}

func (s *AcademicYearService) GetAllAcademicYears() (*[]models.AcademicYear, error) {
	result, err := s.Repo.FindAcademicYears()
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *AcademicYearService) GetAcademicYear(uuid string) (*models.AcademicYear, error) {
	result, err := s.Repo.FindAcademicYearBy("uuid", uuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("Lab Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *AcademicYearService) UpdateAcademicYear(uuid string, req *request.AcademicYear) error {

	result, err := s.GetAcademicYear(uuid)
	if err != nil {
		return err
	}

	model := models.AcademicYear{
		ID:       result.ID,
		FirstDate:      req.FirstDate,
		MiddleDate:     req.MiddleDate,
		MiddleLastDate: req.MiddleLastDate,
		LastDate:       req.LastDate,
		FirstDays:      req.FirstDays,
		MiddleDays:     req.MiddleDays,
		MiddleLastDays: req.MiddleLastDays,
		LastDays:       req.LastDays,
		Semester: req.Semester,
		Year:     req.Year,
	}

	if err := s.Repo.UpdateAcademicYear(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *AcademicYearService) UpdateTimeMonev(uuid string, req *request.UpdateTimeMonev) error {
	result, err := s.GetAcademicYear(uuid)
	if err != nil {
		return err
	}

	var first, middle, middle_last, last string
	if req.FirstDate != "" && req.FirstAmount != "" {
		first = utils.AddDate(req.FirstDate, req.FirstAmount)
	}

	if req.MiddleDate != "" && req.MiddleAmount != "" {
		middle = utils.AddDate(req.MiddleDate, req.MiddleAmount)
	}

	if req.MiddleLastDate != "" && req.MiddleLastAmount != "" {
		middle_last = utils.AddDate(req.MiddleLastDate, req.MiddleLastAmount)
	}

	if req.LastDate != "" && req.LastAmount != "" {
		last = utils.AddDate(req.LastDate, req.LastAmount)
	}

	model := models.AcademicYear{
		ID:             result.ID,
		FirstDate:      first,
		MiddleDate:     middle,
		MiddleLastDate: middle_last,
		LastDate:       last,
	}

	if err := s.Repo.UpdateAcademicYear(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}
	return nil
}

func (s *AcademicYearService) DeleteAcademicYear(uuid string) error {
	result, err := s.GetAcademicYear(uuid)
	if err != nil {
		return err
	}

	model := models.AcademicYear{
		ID: result.ID,
	}

	if err := s.Repo.DeleteAcademicYear(&model); err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return response.VIOLATED_ERR
		}
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *AcademicYearService) UpdateOne(uuid, column string, value interface{}) error {
	result, err := s.GetAcademicYear(uuid)
	if err != nil {
		return err
	}

	if err := s.Repo.UpdateOne(result.ID, column, value); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil
}
