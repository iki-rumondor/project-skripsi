package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/models"
)

type PdfHandlerInterface interface {
	CreateReport(*gin.Context)
}

type PdfServiceInterface interface {
	RpsReport(departmentUuid, yearUuid string) ([]byte, error)
	ModuleReport(departmentUuid, yearUuid string) ([]byte, error)
	ToolReport(departmentUuid, yearUuid string) ([]byte, error)
	SkillReport(departmentUuid, yearUuid string) ([]byte, error)
	FacilitiesReport(departmentUuid, yearUuid string) ([]byte, error)
	TeacherAttendencesReport(departmentUuid, yearUuid string) ([]byte, error)
	StudentAttendencesReport(departmentUuid, yearUuid string) ([]byte, error)
	PlanMatchReport(departmentUuid, yearUuid string) ([]byte, error)
	FinalStudentReport(departmentUuid, yearUuid string) ([]byte, error)
	PassedStudentReport(departmentUuid, yearUuid string) ([]byte, error)
	GradeReport(departmentUuid, yearUuid string) ([]byte, error)
}

type PdfRepoInterface interface {
	LaravelPdf(data interface{}, endpoint string) (*http.Response, error)
	FindDepartmentBy(column string, value interface{}) (*models.Department, error)
	FindAcademicYearBy(column string, value interface{}) (*models.AcademicYear, error)
	FindAcademicPlans(departmentID, yearID uint) (*[]models.AcademicPlan, error)
	FindModules(departmentID, yearID uint) (*[]models.PracticalModule, error)
	FindTools(departmentID, yearID uint) (*[]models.PracticalTool, error)
	FindSkills(departmentID, yearID uint) (*[]models.TeacherSkill, error)
	FindFacilities(departmentID, yearID uint) (*[]models.FacilityCondition, error)
	FindTeacherAttendences(departmentID, yearID uint) (*[]models.TeacherAttendence, error)
	FindStudentAttendences(departmentID, yearID uint) (*[]models.StudentAttendence, error)
}
