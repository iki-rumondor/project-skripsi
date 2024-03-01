package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/response"
	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/utils"
)

type PdfHandler struct {
	Service interfaces.PdfServiceInterface
}

func NewPdfHandler(service interfaces.PdfServiceInterface) interfaces.PdfHandlerInterface {
	return &PdfHandler{
		Service: service,
	}
}

func (h *PdfHandler) CreateReport(c *gin.Context) {
	departmentUuid := c.Param("departmentUuid")
	yearUuid := c.Param("yearUuid")
	typeReport := c.Param("typeReport")

	var dataPDF []byte
	var err error

	switch typeReport {
	case "plans":
		dataPDF, err = h.Service.RpsReport(departmentUuid, yearUuid)
	case "modules":
		dataPDF, err = h.Service.ModuleReport(departmentUuid, yearUuid)
	case "tools":
		dataPDF, err = h.Service.ToolReport(departmentUuid, yearUuid)
	case "skills":
		dataPDF, err = h.Service.SkillReport(departmentUuid, yearUuid)
	case "facilities":
		dataPDF, err = h.Service.FacilitiesReport(departmentUuid, yearUuid)
	case "teacher-attendences":
		dataPDF, err = h.Service.TeacherAttendencesReport(departmentUuid, yearUuid)
	case "student-attendences":
		dataPDF, err = h.Service.StudentAttendencesReport(departmentUuid, yearUuid)
	case "plans-matching":
		dataPDF, err = h.Service.PlanMatchReport(departmentUuid, yearUuid)
	case "final":
		dataPDF, err = h.Service.FinalStudentReport(departmentUuid, yearUuid)
	case "passed":
		dataPDF, err = h.Service.PassedStudentReport(departmentUuid, yearUuid)
	case "grade":
		dataPDF, err = h.Service.GradeReport(departmentUuid, yearUuid)
	default:
		utils.HandleError(c, response.BADREQ_ERR("Endpoint Tidak Ditemukan"))
		return
	}

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=reporting.pdf")
	c.Data(http.StatusOK, "application/pdf", dataPDF)
}
