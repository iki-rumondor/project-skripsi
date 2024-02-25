package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/models"
)

type AcademicPlanHandlerInterface interface {
	CreateAcademicPlan(*gin.Context)
	GetAllAcademicPlans(*gin.Context)
	GetAcademicPlan(*gin.Context)
	UpdateAcademicPlan(*gin.Context)
	DeleteAcademicPlan(*gin.Context)

	GetDepartment(*gin.Context)
	GetMiddle(*gin.Context)
	UpdateMiddle(*gin.Context)
	UpdateLast(*gin.Context)
}

type AcademicPlanServiceInterface interface {
	CreateAcademicPlan(userUuid string, req *request.AcademicPlan) error
	GetAllAcademicPlans(userUuid, yearUuid string) (*[]models.AcademicPlan, error)
	GetAcademicPlan(userUuid, uuid string) (*models.AcademicPlan, error)
	UpdateAcademicPlan(userUuid, uuid string, req *request.UpdateAcademicPlan) error
	DeleteAcademicPlan(userUuid, uuid string) error

	GetUser(userUuid string) (*models.User, error)
	GetAcademicYear(yearUuid string) (*models.AcademicYear, error)
	GetDepartment(departmentUuid, yearUuid string) (*[]models.AcademicPlan, error)
	GetMiddle(userUuid, yearUuid string) (*[]models.AcademicPlan, error)
	UpdateOne(userUuid, uuid, column string, value interface{}) error
}

type AcademicPlanRepoInterface interface {
	FindAcademicPlans(userUuid string, yearID uint) (*[]models.AcademicPlan, error)
	FindUserAcademicPlan(userUuid, uuid string) (*models.AcademicPlan, error)
	FindSubjectBy(column string, value interface{}) (*models.Subject, error)
	CreateAcademicPlan(*models.AcademicPlan) error
	UpdateAcademicPlan(*models.AcademicPlan) error
	DeleteAcademicPlan(*models.AcademicPlan) error
	FindAcademicYearBy(column string, value interface{}) (*models.AcademicYear, error)
	
	FindDepartmentBy(column string, value interface{}) (*models.Department, error)
	FindBy(departmentID, yearID uint, column string, value interface{}) (*[]models.AcademicPlan, error)
	UpdateOne(id uint, column string, value interface{}) error
	FindUser(userUuid string) (*models.User, error)
}
