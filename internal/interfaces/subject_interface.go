package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/models"
)

type SubjectHandlerInterface interface {
	CreateSubject(*gin.Context)
	GetAllSubjects(*gin.Context)
	GetSubject(*gin.Context)
	UpdateSubject(*gin.Context)
	DeleteSubject(*gin.Context)
}

type SubjectServiceInterface interface {
	CreateSubject(*request.Subject) error
	GetAllSubjects() (*[]models.Subject, error)
	GetSubject(string) (*models.Subject, error)
	UpdateSubject(string, *request.Subject) error
	DeleteSubject(string) error
}

type SubjectRepoInterface interface {
	FindSubjects() (*[]models.Subject, error)
	FindSubjectBy(column string, value interface{}) (*models.Subject, error)
	FindUserBy(column string, value interface{}) (*models.User, error)
	CreateSubject(*models.Subject) error
	UpdateSubject(*models.Subject) error
	DeleteSubject(*models.Subject) error
}
