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

type SubjectService struct {
	Repo interfaces.SubjectRepoInterface
}

func NewSubjectService(repo interfaces.SubjectRepoInterface) interfaces.SubjectServiceInterface {
	return &SubjectService{
		Repo: repo,
	}
}

func (s *SubjectService) CreateSubject(userUuid string, req *request.Subject) error {

	user, err := s.Repo.FindUserBy("uuid", userUuid)
	if err != nil {
		return err
	}

	subject := models.Subject{
		DepartmentID: user.Department.ID,
		Name:         req.Name,
		Practical:    req.Practical,
		Code:         req.Code,
	}

	if err := s.Repo.CreateSubject(&subject); err != nil {
		log.Println(err.Error())
		if ok := utils.IsErrorType(err); ok {
			return err
		}
		return response.SERVICE_INTERR
	}
	return nil
}

func (s *SubjectService) GetAllSubjects(userUuid string) (*[]models.Subject, error) {

	subjects, err := s.Repo.FindSubjects(userUuid)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return subjects, nil
}

func (s *SubjectService) GetSubjectsByPlanYear(userUuid, yearUuid string) (*[]models.Subject, error) {

	subjects, err := s.Repo.FindSubjectsByPlanYear(userUuid, yearUuid)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return subjects, nil
}

func (s *SubjectService) GetSubject(userUuid, uuid string) (*models.Subject, error) {
	subject, err := s.Repo.FindUserSubject(userUuid, uuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("Mata Kuliah Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return subject, nil
}

func (s *SubjectService) UpdateSubject(userUuid, uuid string, req *request.Subject) error {

	subject, err := s.GetSubject(userUuid, uuid)
	if err != nil {
		return err
	}

	model := models.Subject{
		ID:           subject.ID,
		DepartmentID: subject.DepartmentID,
		Name:         req.Name,
		Practical:    req.Practical,
		Code:         req.Code,
	}

	if err := s.Repo.UpdateSubject(&model); err != nil {
		log.Println(err.Error())
		if ok := utils.IsErrorType(err); ok {
			return err
		}
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *SubjectService) DeleteSubject(userUuid, uuid string) error {
	subject, err := s.GetSubject(userUuid, uuid)
	if err != nil {
		return err
	}

	model := models.Subject{
		ID: subject.ID,
	}

	if err := s.Repo.DeleteSubject(&model); err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return response.VIOLATED_ERR
		}
		return response.SERVICE_INTERR
	}

	return nil
}
