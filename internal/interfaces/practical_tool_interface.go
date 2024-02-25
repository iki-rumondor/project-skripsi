package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/models"
)

type PracticalToolHandlerInterface interface {
	CreatePracticalTool(*gin.Context)
	GetAllPracticalTools(*gin.Context)
	GetPracticalTool(*gin.Context)
	UpdatePracticalTool(*gin.Context)
	DeletePracticalTool(*gin.Context)

	GetByDepartment(*gin.Context)
}

type PracticalToolServiceInterface interface {
	CreatePracticalTool(userUuid string, req *request.PracticalTool) error
	GetAllPracticalTools(userUuid, yearUuid string) (*[]models.PracticalTool, error)
	GetPracticalTool(userUuid, uuid string) (*models.PracticalTool, error)
	UpdatePracticalTool(userUuid, uuid string, req *request.UpdatePracticalTool) error
	DeletePracticalTool(userUuid, uuid string) error

	GetByDepartment(departmentUuid, yearUuid string) (*[]models.PracticalTool, error)
}

type PracticalToolRepoInterface interface {
	FindPracticalTools(userUuid string, yearID uint) (*[]models.PracticalTool, error)
	FindUserPracticalTool(userUuid, uuid string) (*models.PracticalTool, error)
	FindSubjectBy(column string, value interface{}) (*models.Subject, error)
	FindAcademicYearBy(column string, value interface{}) (*models.AcademicYear, error)
	CreatePracticalTool(*models.PracticalTool) error
	UpdatePracticalTool(*models.PracticalTool) error
	DeletePracticalTool(*models.PracticalTool) error

	FindDepartment(uuid string) (*models.Department, error)
	FindByDepartment(departmentID, yearID uint) (*[]models.PracticalTool, error)
}
