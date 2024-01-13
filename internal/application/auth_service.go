package application

import (
	"errors"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
)

type AuthService struct {
	Repo *repository.Repositories
}

func NewAuthService(repo *repository.Repositories) *AuthService {
	return &AuthService{
		Repo: repo,
	}
}

func (s *AuthService) CreateUser(request *request.Register) error {
	user := &domain.User{
		Username: request.Username,
		Password: request.Password,
		Role:     request.Role,
	}

	if err := s.Repo.AuthRepo.SaveUser(user); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) VerifyUser(user *domain.User) (string, error) {

	result, err := s.Repo.AuthRepo.FindByUsername(user.Username)
	if err != nil {
		return "", errors.New("wrong")
	}

	if err := utils.ComparePassword(result.Password, user.Password); err != nil {
		return "", errors.New("wrong")
	}

	data := map[string]interface{}{
		"id":   result.ID,
		"role": result.Role,
	}

	jwt, err := utils.GenerateToken(data)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func (s *AuthService) GetUsers() (*[]domain.User, error) {

	users, err := s.Repo.AuthRepo.FindUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}
