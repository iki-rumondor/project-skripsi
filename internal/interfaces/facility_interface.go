package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/models"
)

type FacilityHandlerInterface interface {
	CreateFacility(*gin.Context)
	GetAllFacilities(*gin.Context)
	GetFacility(*gin.Context)
	UpdateFacility(*gin.Context)
	DeleteFacility(*gin.Context)
}

type FacilityServiceInterface interface {
	CreateFacility(userUuid string, req *request.Facility) error
	GetAllFacilities(userUuid string) (*[]models.Facility, error)
	GetFacility(userUuid, uuid string) (*models.Facility, error)
	UpdateFacility(userUuid, uuid string, req *request.Facility) error
	DeleteFacility(userUuid, uuid string) error
}

type FacilityRepoInterface interface {
	FindFacilities(userUuid string) (*[]models.Facility, error)
	FindUserBy(column string, value interface{}) (*models.User, error)
	FindUserFacility(userUuid, uuid string) (*models.Facility, error)
	CreateFacility(*models.Facility) error
	UpdateFacility(*models.Facility) error
	DeleteFacility(*models.Facility) error
}
