package services

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/http/response"
	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
)

type AcademicPlanService struct {
	Repo interfaces.AcademicPlanRepoInterface
}

func NewAcademicPlanService(repo interfaces.AcademicPlanRepoInterface) interfaces.AcademicPlanServiceInterface {
	return &AcademicPlanService{
		Repo: repo,
	}
}

func (s *AcademicPlanService) CreateAcademicPlan(userUuid string, req *request.AcademicPlan) error {

	subject, err := s.Repo.FindSubjectBy("uuid", req.SubjectUuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Mata Kuliah Tidak Ditemukan")
		}
		return response.SERVICE_INTERR
	}

	academic_year, err := s.Repo.FindAcademicYearBy("uuid", req.AcademicYearUuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Tahun Ajaran Tidak Ditemukan")
		}
		return response.SERVICE_INTERR
	}

	if !req.Available && req.Note == "" {
		return response.BADREQ_ERR("Keterangan Tidak Valid")
	}

	rps := models.AcademicPlan{
		Available:      &req.Available,
		Note:           req.Note,
		SubjectID:      subject.ID,
		AcademicYearID: academic_year.ID,
	}

	if err := s.Repo.Create(&rps); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}
	return nil
}

func (s *AcademicPlanService) GetAllAcademicPlans(userUuid, yearUuid string) (*[]models.AcademicPlan, error) {

	year, err := s.GetAcademicYear(yearUuid)
	if err != nil {
		return nil, err
	}

	result, err := s.Repo.FindAcademicPlans(userUuid, year.ID)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *AcademicPlanService) GetAllRps(userUuid, yearUuid string) (*[]models.Rps, error) {

	year, err := s.GetAcademicYear(yearUuid)
	if err != nil {
		return nil, err
	}

	result, err := s.Repo.FindRps(userUuid, year.ID)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *AcademicPlanService) GetDepartment(departmentUuid, yearUuid string) (*[]models.AcademicPlan, error) {

	department, err := s.Repo.FindDepartmentBy("uuid", departmentUuid)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	year, err := s.Repo.FindAcademicYearBy("uuid", yearUuid)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	result, err := s.Repo.FindBy(department.ID, year.ID, "academic_year_id", year.ID)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *AcademicPlanService) GetAcademicPlan(userUuid, uuid string) (*models.AcademicPlan, error) {
	result, err := s.Repo.FindUserAcademicPlan(userUuid, uuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("RPS Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *AcademicPlanService) GetRps(userUuid, uuid string) (*models.Rps, error) {
	result, err := s.Repo.FirstRps(userUuid, uuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("RPS Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *AcademicPlanService) GetUser(userUuid string) (*models.User, error) {
	result, err := s.Repo.FindUser(userUuid)

	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("User Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *AcademicPlanService) GetAcademicYear(yearUuid string) (*models.AcademicYear, error) {
	result, err := s.Repo.FindAcademicYearBy("uuid", yearUuid)

	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("Tahun Ajaran Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *AcademicPlanService) GetMiddle(userUuid, yearUuid string) (*[]models.AcademicPlan, error) {
	user, err := s.GetUser(userUuid)
	if err != nil {
		return nil, err
	}

	year, err := s.GetAcademicYear(yearUuid)
	if err != nil {
		return nil, err
	}

	result, err := s.Repo.FindBy(user.Department.ID, year.ID, "available", true)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *AcademicPlanService) UpdateOne(userUuid, uuid, column string, value interface{}) error {
	result, err := s.GetAcademicPlan(userUuid, uuid)
	if err != nil {
		return err
	}

	if err := s.Repo.UpdateOne(result.ID, column, value); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *AcademicPlanService) UpdateAcademicPlan(userUuid, uuid, fileName string, req *request.UpdateAcademicPlan) error {

	result, err := s.GetRps(userUuid, uuid)
	if err != nil {
		return err
	}

	if !req.Status && req.Note == "" {
		return response.BADREQ_ERR("Keterangan Tidak Valid")
	}

	dataUpdate := map[string]interface{}{
		"status":    req.Status,
		"note":      req.Note,
		"file_name": fileName,
	}

	if req.Note == "" || req.Note == "null" {
		dataUpdate["note"] = nil
	}

	if fileName == "" || fileName == "null" {
		dataUpdate["file_name"] = nil
	}

	model := models.Rps{
		ID: result.ID,
		// Status:   req.Status,
		// Note:     req.Note,
		// FileName: fileName,
	}

	if err := s.Repo.Updates(&model, dataUpdate); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	if result.FileName != nil {
		pathFile := filepath.Join("internal/files/rps", *result.FileName)
		if err := os.Remove(pathFile); err != nil {
			log.Println(err.Error())
		}
	}

	return nil
}

func (s *AcademicPlanService) DeleteAcademicPlan(userUuid, uuid string) error {
	result, err := s.GetAcademicPlan(userUuid, uuid)
	if err != nil {
		return err
	}

	model := models.AcademicPlan{
		ID: result.ID,
	}

	if err := s.Repo.DeleteAcademicPlan(&model); err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return response.VIOLATED_ERR
		}
		return response.SERVICE_INTERR
	}

	return nil
}
