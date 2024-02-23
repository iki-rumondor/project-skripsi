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

type UserService struct {
	Repo interfaces.UserRepoInterface
}

func NewUserService(repo interfaces.UserRepoInterface) interfaces.UserServiceInterface {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) GetUser(column string, value interface{}) (*models.User, error) {
	user, err := s.Repo.FindUserBy(column, value)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.NOTFOUND_ERR("User Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return user, nil
}
func (s *UserService) VerifyUser(req *request.SignIn) (string, error) {

	user, err := s.Repo.FindUserBy("username", req.Username)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", &response.Error{
				Code:    401,
				Message: "Username atau Password Salah",
			}
		}
		return "", response.SERVICE_INTERR
	}

	if err := utils.ComparePassword(user.Password, req.Password); err != nil {
		return "", &response.Error{
			Code:    401,
			Message: "Username atau password salah",
		}
	}

	jwt, err := utils.GenerateToken(user.Uuid, user.Role.Name)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func (s *UserService) GetCountSubjects() (*response.SubjectsCount, error) {
	subjects, err := s.Repo.FindSubjects()
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	practicalSubjects, err := s.Repo.FindPracticalSubjects()
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	return &response.SubjectsCount{
		General:   len(*subjects),
		Practical: len(*practicalSubjects),
	}, nil
}

func (s *UserService) CountMonevByYear(userUuid, yearUuid string) (map[string]int, error) {

	user, err := s.GetUser("uuid", userUuid)
	if err != nil {
		return nil, err
	}

	year, err := s.Repo.GetOne("academic_years", "uuid", yearUuid)
	if err != nil {
		return nil, err
	}

	yearID := year["id"].(uint64)

	result, err := s.Repo.CountMonevByYear(user.Department.ID, uint(yearID))
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *UserService) CountDepartmentMonev(departmentUuid, yearUuid string) (map[string]int, error) {

	department, err := s.Repo.GetOne("departments", "uuid", departmentUuid)
	if err != nil {
		return nil, err
	}

	departmentID := department["id"].(uint64)

	year, err := s.Repo.GetOne("academic_years", "uuid", yearUuid)
	if err != nil {
		return nil, err
	}

	yearID := year["id"].(uint64)

	result, err := s.Repo.CountMonevByYear(uint(departmentID), uint(yearID))
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *UserService) Update(id uint, tableName, column string, value interface{}) error {

	if err := s.Repo.Update(id, tableName, column, value); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *UserService) GetAll(tableName string) ([]map[string]interface{}, error) {
	result, err := s.Repo.GetAll(tableName)
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}
