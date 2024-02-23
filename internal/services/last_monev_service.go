package services

import (
	"errors"
	"log"

	"github.com/iki-rumondor/go-monev/internal/http/response"
	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
)

type LastMonevService struct {
	Repo interfaces.LastMonevRepoInterface
}

func NewLastMonevService(repo interfaces.LastMonevRepoInterface) interfaces.LastMonevServiceInterface {
	return &LastMonevService{
		Repo: repo,
	}
}

func (s *LastMonevService) CountLastMonev(userUuid, yearUuid string) (map[string]int, error) {
	year, err := s.GetAcademicYear(yearUuid)
	if err != nil {
		return nil, err
	}

	user, err := s.GetUser(userUuid)
	if err != nil {
		return nil, err
	}

	result, err := s.Repo.CountLastMonev(user.Department.ID, year.ID)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *LastMonevService) GetAcademicYear(yearUuid string) (*models.AcademicYear, error) {
	result, err := s.Repo.FindAcademicYear(yearUuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("Tahun Ajaran Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *LastMonevService) GetUser(userUuid string) (*models.User, error) {
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

func (s *LastMonevService) GetSubject(userUuid, uuid string) (*models.Subject, error) {
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

func (s *LastMonevService) GetStudentAttendence(userUuid, uuid string) (*models.StudentAttendence, error) {
	result, err := s.Repo.FindStudentAttendence(userUuid, uuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("Kelulusan Mahasiswa Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *LastMonevService) GetTeacherAttendence(userUuid, uuid string) (*models.TeacherAttendence, error) {
	result, err := s.Repo.FindTeacherAttendence(userUuid, uuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("Kelulusan Mahasiswa Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *LastMonevService) UpdateStudentAttendence(userUuid, uuid, column string, value interface{}) error {

	result, err := s.GetStudentAttendence(userUuid, uuid)
	if err != nil {
		return err
	}

	if err := s.Repo.UpdateStudentAttendence(result, column, value); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *LastMonevService) UpdateTeacherAttendence(userUuid, uuid, column string, value interface{}) error {

	result, err := s.GetTeacherAttendence(userUuid, uuid)
	if err != nil {
		return err
	}

	if err := s.Repo.UpdateTeacherAttendence(result, column, value); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil
}

