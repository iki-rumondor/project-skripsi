package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/models"
)

type TeacherSkillHandlerInterface interface {
	CreateTeacherSkill(*gin.Context)
	GetByYear(*gin.Context)
	GetAllTeacherSkills(*gin.Context)
	GetTeacherSkill(*gin.Context)
	UpdateTeacherSkill(*gin.Context)
	DeleteTeacherSkill(*gin.Context)

	GetByDepartment(*gin.Context)
}

type TeacherSkillServiceInterface interface {
	CreateTeacherSkill(userUuid string, req *request.TeacherSkill) error
	GetAllTeacherSkills(userUuid string) (*[]models.TeacherSkill, error)
	GetTeacherSkillsByYear(userUuid, yearUuid string) (*[]models.TeacherSkill, error)
	GetTeacherSkill(userUuid, uuid string) (*models.TeacherSkill, error)
	UpdateTeacherSkill(userUuid, uuid string, req *request.UpdateTeacherSkill) error
	DeleteTeacherSkill(userUuid, uuid string) error

	GetByDepartment(departmentUuid, yearUuid string) (*[]models.TeacherSkill, error)
}

type TeacherSkillRepoInterface interface {
	FindTeacherSkills(userUuid string) (*[]models.TeacherSkill, error)
	FindTeacherSkillsByYear(userUuid string, yearID uint) (*[]models.TeacherSkill, error)
	FindUserTeacherSkill(userUuid, uuid string) (*models.TeacherSkill, error)
	FindTeacherBy(column string, value interface{}) (*models.Teacher, error)
	FindSubjectByUuid(uuid string) (*models.Subject, error)
	FindAcademicYearByUuid(uuid string) (*models.AcademicYear, error)
	CreateTeacherSkill(*models.TeacherSkill) error
	UpdateTeacherSkill(*models.TeacherSkill) error
	DeleteTeacherSkill(*models.TeacherSkill) error

	FindAcademicYearBy(column string, value interface{}) (*models.AcademicYear, error)
	FindDepartment(uuid string) (*models.Department, error)
	FindByDepartment(departmentID, yearID uint) (*[]models.TeacherSkill, error)
}
