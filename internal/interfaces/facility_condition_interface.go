package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/models"
)

type FacilityConditionHandlerInterface interface {
	CreateFacilityCondition(*gin.Context)
	GetFacilityConditionsByYear(*gin.Context)
	GetFacilityOptions(*gin.Context)
	GetFacilityCondition(*gin.Context)
	UpdateFacilityCondition(*gin.Context)
	DeleteFacilityCondition(*gin.Context)

	GetByDepartment(*gin.Context)
}

type FacilityConditionServiceInterface interface {
	CreateFacilityCondition(userUuid string, req *request.FacilityCondition) error
	GetFacilityConditionsByYear(userUuid, yearUuid string) (*[]models.FacilityCondition, error)
	GetFacilityOptions(userUuid, yearUuid string) (*[]models.Facility, error)
	GetFacilityCondition(userUuid, uuid string) (*models.FacilityCondition, error)
	UpdateFacilityCondition(userUuid, uuid string, req *request.UpdateFacilityCondition) error
	DeleteFacilityCondition(userUuid, uuid string) error
	
	GetByDepartment(departmentUuid, yearUuid string) (*[]models.FacilityCondition, error)
}

type FacilityConditionRepoInterface interface {
	FindFacilityConditionsByYear(userUuid string, yearID uint) (*[]models.FacilityCondition, error)
	FindUserFacilityCondition(userUuid, uuid string) (*models.FacilityCondition, error)
	FindFacilityOptions(userUuid string, yearID uint) (*[]models.Facility, error)
	FindFacilityBy(column string, value interface{}) (*models.Facility, error)
	FindAcademicYearBy(column string, value interface{}) (*models.AcademicYear, error)
	CreateFacilityCondition(*models.FacilityCondition) error
	UpdateFacilityCondition(*models.FacilityCondition) error
	DeleteFacilityCondition(*models.FacilityCondition) error

	FindDepartment(uuid string) (*models.Department, error)
	FindByDepartment(departmentID, yearID uint) (*[]models.FacilityCondition, error)
}
