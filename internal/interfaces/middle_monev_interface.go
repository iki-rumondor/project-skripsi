package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/models"
)

type MiddleMonevHandlerInterface interface {
	CreateTeacherAttendence(*gin.Context)
	CreateStudentAttendence(*gin.Context)
	GetTeacherAttendences(*gin.Context)
	GetStudentAttendences(*gin.Context)
	CountTotalMonev(*gin.Context)
	UpdateLastTeacherAttendence(*gin.Context)
	UpdateLastStudentAttendence(*gin.Context)
	DeleteTeacherAttendence(*gin.Context)
	DeleteStudentAttendence(*gin.Context)
}

type MiddleMonevServiceInterface interface {
	CreateTeacherAttendence(userUuid string, req *request.CreateTeacherAttendence) error
	CreateStudentAttendence(userUuid string, req *request.CreateStudentAttendence) error
	GetTeacherAttendences(userUuid, yearUuid string) (*[]models.TeacherAttendence, error)
	GetStudentAttendences(userUuid, yearUuid string) (*[]models.StudentAttendence, error)
	GetTeacherAttendence(userUuid, uuid string) (*models.TeacherAttendence, error)
	GetStudentAttendence(userUuid, uuid string) (*models.StudentAttendence, error)
	CountTotalMonev(userUuid, yearUuid string) (map[string]int, error)
	UpdateTeacherAttendence(userUuid, uuid string, model *models.TeacherAttendence) error
	UpdateStudentAttendence(userUuid, uuid string, model *models.StudentAttendence) error
	DeleteTeacherAttendence(userUuid, uuid string) error
	DeleteStudentAttendence(userUuid, uuid string) error
}

type MiddleMonevRepoInterface interface {
	CreateTeacherAttendence(*models.TeacherAttendence) error
	CreateStudentAttendence(*models.StudentAttendence) error
	FindTeacherAttendences(userUuid string, yearID uint) (*[]models.TeacherAttendence, error)
	FindStudentAttendences(userUuid string, yearID uint) (*[]models.StudentAttendence, error)
	FindTeacherAttendence(userUuid, uuid string) (*models.TeacherAttendence, error)
	FindStudentAttendence(userUuid, uuid string) (*models.StudentAttendence, error)
	FindUserSubject(userUuid, uuid string) (*models.Subject, error)
	FindAcademicYear(uuid string) (*models.AcademicYear, error)
	CountTotalMonev(userUuid string, yearID uint) (map[string]int, error)
	UpdateTeacherAttendence(model *models.TeacherAttendence) error
	UpdateStudentAttendence(model *models.StudentAttendence) error
	DeleteTeacherAttendence(*models.TeacherAttendence) error
	DeleteStudentAttendence(*models.StudentAttendence) error
}
