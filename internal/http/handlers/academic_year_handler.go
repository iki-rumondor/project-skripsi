package handlers

import (
	"fmt"
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
		var tools = []response.PracticalTool{}
		for _, item := range *item.PracticalTools {
			tools = append(tools, response.PracticalTool{
				Uuid: item.Uuid,
			})
		}

		var plans = []response.AcademicPlan{}
		for _, item := range *item.AcademicPlans {
			plans = append(plans, response.AcademicPlan{
				Uuid: item.Uuid,
			})
		}

		var modules = []response.PracticalModule{}
		for _, item := range *item.PracticalModules {
			modules = append(modules, response.PracticalModule{
				Uuid: item.Uuid,
			})
		}

		yearName := fmt.Sprintf("%s %s", item.Semester, item.Year)

		first := utils.AddDate(item.FirstDate, item.FirstDays)
		middle := utils.AddDate(item.MiddleDate, item.MiddleDays)
		middle_last := utils.AddDate(item.MiddleLastDate, item.MiddleLastDays)
		last := utils.AddDate(item.LastDate, item.LastDays)

		status := "0"

		if ok := utils.IsNowInRange(first); ok {
			status = "1"
		}

		if ok := utils.IsNowInRange(middle); ok {
			status = "2"
		}

		if ok := utils.IsNowInRange(middle_last); ok {
			status = "3"
		}

		if ok := utils.IsNowInRange(last); ok {
			status = "4"
		}

		resp = append(resp, &response.AcademicYear{
			Uuid:            item.Uuid,
			Name:            yearName,
			Semester:        item.Semester,
			Year:            item.Year,
			FirstDate:       item.FirstDate,
			MiddleDate:      item.MiddleDate,
			MiddleLastDate:  item.MiddleLastDate,
			LastDate:        item.LastDate,
			FirstDays:       item.FirstDays,
			MiddleDays:      item.MiddleDays,
			MiddleLastDays:  item.MiddleLastDays,
			LastDays:        item.LastDays,
			FirstRange:      first,
			MiddleRange:     middle,
			MiddleLastRange: middle_last,
			LastRange:       last,
			Status:          status,
			AcademicPlan:    &plans,
			PracticalTool:   &tools,
			PracticalModule: &modules,
			CreatedAt:       item.CreatedAt,
			UpdatedAt:       item.UpdatedAt,
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
	yearName := fmt.Sprintf("%s %s", result.Semester, result.Year)

	first := utils.AddDate(result.FirstDate, result.FirstDays)
	middle := utils.AddDate(result.MiddleDate, result.MiddleDays)
	middle_last := utils.AddDate(result.MiddleLastDate, result.MiddleLastDays)
	last := utils.AddDate(result.LastDate, result.LastDays)

	status := "0"

	if ok := utils.IsNowInRange(first); ok {
		status = "1"
	}

	if ok := utils.IsNowInRange(middle); ok {
		status = "2"
	}

	if ok := utils.IsNowInRange(middle_last); ok {
		status = "3"
	}

	if ok := utils.IsNowInRange(last); ok {
		status = "4"
	}

	resp := &response.AcademicYear{
		Uuid:            result.Uuid,
		Semester:        result.Semester,
		Year:            result.Year,
		Status:          status,
		FirstDate:       result.FirstDate,
		MiddleDate:      result.MiddleDate,
		MiddleLastDate:  result.MiddleLastDate,
		LastDate:        result.LastDate,
		FirstDays:       result.FirstDays,
		MiddleDays:      result.MiddleDays,
		MiddleLastDays:  result.MiddleLastDays,
		LastDays:        result.LastDays,
		FirstRange:      first,
		MiddleRange:     middle,
		MiddleLastRange: middle_last,
		LastRange:       last,
		Name:            yearName,
		CreatedAt:       result.CreatedAt,
		UpdatedAt:       result.UpdatedAt,
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

func (h *AcademicYearHandler) UpdateTimeMonev(c *gin.Context) {
	var body request.UpdateTimeMonev
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	uuid := c.Param("uuid")
	if err := h.Service.UpdateTimeMonev(uuid, &body); err != nil {
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

func (h *AcademicYearHandler) UpdateOpen(c *gin.Context) {
	var body request.AcademicYearOpen
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	uuid := c.Param("uuid")
	if err := h.Service.UpdateOne(uuid, "open", body.Open); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Tahun Ajaran Berhasil Diperbarui"))
}
