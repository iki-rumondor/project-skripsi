package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/models"
)

type SubjectHandlerInterface interface {
	CreateSubject(*gin.Context)
	GetAllSubjects(*gin.Context)
	GetAllPracticalSubjects(*gin.Context)
	GetSubject(*gin.Context)
	UpdateSubject(*gin.Context)
	DeleteSubject(*gin.Context)
	GetOuterSubjects(c *gin.Context)
	GetTeacherAttendenceSubjects(*gin.Context)
	GetStudentAttendenceSubjects(*gin.Context)
}

type SubjectServiceInterface interface {
	CreateSubject(userUuid string, req *request.Subject) error
	GetAllSubjects(userUuid string) (*[]models.Subject, error)
	GetPracticalSubjects(userUuid string) (*[]models.Subject, error)
	GetTeacherAttendenceSubjects(userUuid, yearUuid string) (*[]models.Subject, error)
	GetStudentAttendenceSubjects(userUuid, yearUuid string) (*[]models.Subject, error)
	GetOuterSubjects(userUuid, yearUuid, tableName string) (*[]models.Subject, error)
	GetSubject(userUuid, uuid string) (*models.Subject, error)
	UpdateSubject(userUuid, uuid string, req *request.UpdateSubject) error
	DeleteSubject(userUuid, uuid string) error
}

type SubjectRepoInterface interface {
	FindSubjects(userUuid string) (*[]models.Subject, error)
	FindPracticalSubjects(userUuid string) (*[]models.Subject, error)
	FindUserBy(column string, value interface{}) (*models.User, error)
	FindUserSubject(userUuid, uuid string) (*models.Subject, error)
	FindTeacherAttendenceSubjects(userUuid string, yearID uint) (*[]models.Subject, error)
	FindStudentAttendenceSubjects(userUuid string, yearID uint) (*[]models.Subject, error)
	FindOuterSubjects(departmentID, yearID uint, tableName string) (*[]models.Subject, error)
	FindAcademicYear(uuid string) (*models.AcademicYear, error)
	CreateSubject(*models.Subject) error
	UpdateSubject(*models.Subject) error
	DeleteSubject(*models.Subject) error
}
