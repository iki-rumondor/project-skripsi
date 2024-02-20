package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/models"
)

type LastMonevHandlerInterface interface {
	CountLastMonev(*gin.Context)
	UpdateStudentPass(*gin.Context)
	UpdateStudentFinal(*gin.Context)
	UpdateTeacherGrade(*gin.Context)
	GetAllUserStudentPassed(*gin.Context)
	DeleteStudentPassed(*gin.Context)
}

type LastMonevServiceInterface interface {
	CountLastMonev(userUuid, yearUuid string) (map[string]int, error)
	UpdateStudentAttendence(userUuid, uuid, column string, value interface{}) error
	UpdateTeacherAttendence(userUuid, uuid, column string, value interface{}) error
	GetAllUserStudentPassed(userUuid, yearUuid string) (*[]models.StudentPassed, error)
	DeleteStudentPassed(userUuid, uuid string) error
}

type LastMonevRepoInterface interface {
	CountLastMonev(departmentID, yearID uint) (map[string]int, error)
	UpdateStudentAttendence(model *models.StudentAttendence, column string, value interface{}) error
	UpdateTeacherAttendence(model *models.TeacherAttendence, column string, value interface{}) error
	FindAcademicYear(uuid string) (*models.AcademicYear, error)
	FindUser(uuid string) (*models.User, error)
	FindUserSubject(userUuid, uuid string) (*models.Subject, error)
	FindStudentAttendence(userUuid, uuid string) (*models.StudentAttendence, error)
	FindTeacherAttendence(userUuid, uuid string) (*models.TeacherAttendence, error)
	FindAllStudentPassed(departmentID, yearID uint) (*[]models.StudentPassed, error)
	FindStudentPassed(departmentID uint, uuid string) (*models.StudentPassed, error)
	DeleteStudentPassed(model *models.StudentPassed) error
}
