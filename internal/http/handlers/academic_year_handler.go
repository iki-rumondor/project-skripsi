package handlers

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/http/response"
	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/utils"
)

type AcademicYearHandler struct {
	Service interfaces.AcademicYearServiceInterface
}

func NewAcademicYearHandler(service interfaces.AcademicYearServiceInterface) interfaces.AcademicYearHandlerInterface {
	return &AcademicYearHandler{
		Service: service,
	}
}

func (h *AcademicYearHandler) CreateAcademicYear(c *gin.Context) {
	var body request.AcademicYear
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if err := h.Service.CreateAcademicYear(&body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Tahun Ajaran Berhasil Ditambahkan"))
}

func (h *AcademicYearHandler) GetAllAcademicYears(c *gin.Context) {
	result, err := h.Service.GetAllAcademicYears()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var resp []*response.AcademicYear
	for _, item := range *result {
		resp = append(resp, &response.AcademicYear{
			Uuid:      item.Uuid,
			Name:      item.Name,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *AcademicYearHandler) GetAcademicYear(c *gin.Context) {
	uuid := c.Param("uuid")
	result, err := h.Service.GetAcademicYear(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	resp := &response.AcademicYear{
		Uuid:      result.Uuid,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *AcademicYearHandler) UpdateAcademicYear(c *gin.Context) {
	var body request.AcademicYear
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	uuid := c.Param("uuid")
	if err := h.Service.UpdateAcademicYear(uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response.SUCCESS_RES("Tahun Ajaran Berhasil Diperbarui"))
}

func (h *AcademicYearHandler) DeleteAcademicYear(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := h.Service.DeleteAcademicYear(uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Tahun Ajaran Berhasil Dihapus"))
}
