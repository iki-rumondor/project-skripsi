package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/models"
)

type AcademicYearHandlerInterface interface {
	CreateAcademicYear(*gin.Context)
	GetAllAcademicYears(*gin.Context)
	GetAcademicYear(*gin.Context)
	UpdateAcademicYear(*gin.Context)
	UpdateOpen(*gin.Context)
	DeleteAcademicYear(*gin.Context)	
	UpdateTimeMonev(*gin.Context)
}

type AcademicYearServiceInterface interface {
	CreateAcademicYear(*request.AcademicYear) error
	GetAllAcademicYears() (*[]models.AcademicYear, error)
	GetAcademicYear(string) (*models.AcademicYear, error)
	UpdateAcademicYear(string, *request.AcademicYear) error
	UpdateTimeMonev(string, *request.UpdateTimeMonev) error
	DeleteAcademicYear(string) error
	UpdateOne(uuid, column string, value interface{}) error
}

type AcademicYearRepoInterface interface {
	FindAcademicYears() (*[]models.AcademicYear, error)
	FindAcademicYearBy(column string, value interface{}) (*models.AcademicYear, error)
	CreateAcademicYear(*models.AcademicYear) error
	UpdateAcademicYear(*models.AcademicYear) error
	DeleteAcademicYear(*models.AcademicYear) error
	UpdateOne(id uint, column string, value interface{}) error
}
