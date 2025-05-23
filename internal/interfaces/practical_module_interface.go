package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/models"
)

type PracticalModuleHandlerInterface interface {
	CreatePracticalModule(*gin.Context)
	GetAllPracticalModules(*gin.Context)
	GetPracticalModule(*gin.Context)
	UpdatePracticalModule(*gin.Context)
	DeletePracticalModule(*gin.Context)
	GetByDepartment(*gin.Context)
}

type PracticalModuleServiceInterface interface {
	CreatePracticalModule(userUuid string, req *request.PracticalModule) error
	GetAllPracticalModules(userUuid, yearUuid string) (*[]models.PracticalModule, error)
	GetPracticalModule(userUuid, uuid string) (*models.PracticalModule, error)
	UpdatePracticalModule(userUuid, uuid string, req *request.UpdatePracticalModule) error
	DeletePracticalModule(userUuid, uuid string) error
	GetByDepartment(departmentUuid, yearUuid string) (*[]models.PracticalModule, error)
}

type PracticalModuleRepoInterface interface {
	FindPracticalModules(userUuid string, yearID uint) (*[]models.PracticalModule, error)
	FindUserPracticalModule(userUuid, uuid string) (*models.PracticalModule, error)
	FindSubjectBy(column string, value interface{}) (*models.Subject, error)
	FindAcademicYearBy(column string, value interface{}) (*models.AcademicYear, error)
	FindLaboratoryBy(column string, value interface{}) (*models.Laboratory, error)
	CreatePracticalModule(*models.PracticalModule) error
	UpdatePracticalModule(*models.PracticalModule) error
	DeletePracticalModule(*models.PracticalModule) error

	FindDepartment(uuid string) (*models.Department, error)
	FindByDepartment(departmentID, yearID uint) (*[]models.PracticalModule, error)
}
