package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/http/response"
	"github.com/iki-rumondor/go-monev/internal/models"
)

type UserHandlerInterface interface {
	SignIn(*gin.Context)
	GetDashboardAdmin(*gin.Context)
	CountMonevByYear(*gin.Context)
	UpdateStepMonev(*gin.Context)
	GetSettings(*gin.Context)
	GetDepartmentMonev(*gin.Context)
	GetDepartmentData(*gin.Context)
	GetUsers(*gin.Context)
	CreateUser(*gin.Context)
	GetRoles(*gin.Context)
	GetDepartmentsChart(*gin.Context)
	GetCurrentAcademicYear(*gin.Context)
	// GetTeacherAttendanceClasses(*gin.Context)
}

type UserServiceInterface interface {
	VerifyUser(*request.SignIn) (string, error)
	GetDashboardAdmin() (map[string]interface{}, error)
	GetUser(column string, value interface{}) (*models.User, error)
	GetAllUser() (*[]models.User, error)
	CountMonevByYear(userUuid, yearUuid string) (map[string]int, error)
	CountDepartmentMonev(departmentUuid, yearUuid string) (map[string]interface{}, error)
	Update(id uint, tableName, column string, value interface{}) error
	GetAll(tableName string) ([]map[string]interface{}, error)
	CreateUser(req *request.CreateUser) error
	GetDepartmentsChart(yearUuid string) (map[string]interface{}, error)
	GetCurrentAcademicYear() (*response.AcademicYear, error)
	// GetTeacherAttendanceClasses() ([]string, error)
}

type UserRepoInterface interface {
	FindUserBy(column string, value interface{}) (*models.User, error)
	FindUsers() (*[]models.User, error)
	FindSubjects() (*[]models.Subject, error)
	FindPracticalSubjects() (*[]models.Subject, error)
	CountMonevByYear(departmentID, yearID uint) (map[string]int, error)
	GetAll(tableName string) ([]map[string]interface{}, error)
	GetOne(tableName, column string, value interface{}) (map[string]interface{}, error)
	Update(id uint, tableName, column string, value interface{}) error
	CreateUser(model *models.User) error
	First(dest interface{}, condition string) error
	Find(dest interface{}, condition string) error
	FindDepartments(dest *[]models.Department) error
	FirstWithOrder(dest interface{}, condition, order string) error
	FindTeacherSkills(dest *[]models.TeacherSkill, departmentID, yearID uint) error
}
