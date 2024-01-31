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

type SubjectService struct {
	Repo interfaces.SubjectRepoInterface
}

func NewSubjectService(repo interfaces.SubjectRepoInterface) interfaces.SubjectServiceInterface {
	return &SubjectService{
		Repo: repo,
	}
}

func (s *SubjectService) CreateSubject(req *request.Subject) error {
	user, err := s.Repo.FindUserBy("uuid", req.UserUuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("User tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}

	subject := models.Subject{
		DepartmentID: user.Department.ID,
		Name:         req.Name,
		Code:         req.Code,
	}

	if err := s.Repo.CreateSubject(&subject); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}
	return nil
}

func (s *SubjectService) GetAllSubjects() (*[]models.Subject, error) {
	subjects, err := s.Repo.FindSubjects()
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return subjects, nil
}

func (s *SubjectService) GetSubject(uuid string) (*models.Subject, error) {
	subject, err := s.Repo.FindSubjectBy("uuid", uuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("Mata Kuliah Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return subject, nil
}

func (s *SubjectService) UpdateSubject(uuid string, req *request.Subject) error {
	user, err := s.Repo.FindUserBy("uuid", req.UserUuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("User tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}

	subject, err := s.GetSubject(uuid)
	if err != nil {
		return err
	}

	model := models.Subject{
		ID:           subject.ID,
		DepartmentID: user.Department.ID,
		Name:         req.Name,
		Code:         req.Code,
	}

	if err := s.Repo.UpdateSubject(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *SubjectService) DeleteSubject(uuid string) error {
	subject, err := s.GetSubject(uuid)
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
