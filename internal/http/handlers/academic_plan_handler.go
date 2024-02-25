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

type AcademicPlanHandler struct {
	Service interfaces.AcademicPlanServiceInterface
}

func NewAcademicPlanHandler(service interfaces.AcademicPlanServiceInterface) interfaces.AcademicPlanHandlerInterface {
	return &AcademicPlanHandler{
		Service: service,
	}
}

func (h *AcademicPlanHandler) CreateAcademicPlan(c *gin.Context) {
	var body request.AcademicPlan
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	if err := h.Service.CreateAcademicPlan(userUuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("RPS Berhasil Ditambahkan"))
}

func (h *AcademicPlanHandler) GetAllAcademicPlans(c *gin.Context) {
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}
	yearUuid := c.Param("yearUuid")

	result, err := h.Service.GetAllAcademicPlans(userUuid, yearUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var resp []*response.AcademicPlan
	for _, item := range *result {
		yearName := fmt.Sprintf("%s %s", item.AcademicYear.Semester, item.AcademicYear.Year)
		resp = append(resp, &response.AcademicPlan{
			Uuid:      item.Uuid,
			Available: item.Available,
			Note:      item.Note,
			Middle:    item.Middle,
			Last:      item.Last,
			AcademicYear: &response.AcademicYear{
				Uuid: item.AcademicYear.Uuid,
				Name: yearName,
			},
			Subject: &response.Subject{
				Uuid: item.Subject.Uuid,
				Name: item.Subject.Name,
				Code: item.Subject.Code,
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *AcademicPlanHandler) GetDepartment(c *gin.Context) {
	departmentUuid := c.Param("departmentUuid")
	yearUuid := c.Param("yearUuid")

	result, err := h.Service.GetDepartment(departmentUuid, yearUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var resp []*response.AcademicPlan
	for _, item := range *result {
		yearName := fmt.Sprintf("%s %s", item.AcademicYear.Semester, item.AcademicYear.Year)
		resp = append(resp, &response.AcademicPlan{
			Uuid:      item.Uuid,
			Available: item.Available,
			Note:      item.Note,
			Middle:    item.Middle,
			Last:      item.Last,
			AcademicYear: &response.AcademicYear{
				Uuid: item.AcademicYear.Uuid,
				Name: yearName,
			},
			Subject: &response.Subject{
				Uuid: item.Subject.Uuid,
				Name: item.Subject.Name,
				Code: item.Subject.Code,
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *AcademicPlanHandler) GetAcademicPlan(c *gin.Context) {
	uuid := c.Param("uuid")
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}
	result, err := h.Service.GetAcademicPlan(userUuid, uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	yearName := fmt.Sprintf("%s %s", result.AcademicYear.Semester, result.AcademicYear.Year)

	resp := &response.AcademicPlan{
		Uuid:      result.Uuid,
		Available: result.Available,
		Note:      result.Note,
		Middle:    result.Middle,
		Last:      result.Last,
		AcademicYear: &response.AcademicYear{
			Uuid: result.AcademicYear.Uuid,
			Name: yearName,
		},
		Subject: &response.Subject{
			Uuid: result.Subject.Uuid,
			Name: result.Subject.Name,
			Code: result.Subject.Code,
		},
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *AcademicPlanHandler) GetMiddle(c *gin.Context) {
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	yearUuid := c.Param("yearUuid")

	result, err := h.Service.GetMiddle(userUuid, yearUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var resp []*response.AcademicPlan
	for _, item := range *result {
		yearName := fmt.Sprintf("%s %s", item.AcademicYear.Semester, item.AcademicYear.Year)
		resp = append(resp, &response.AcademicPlan{
			Uuid:      item.Uuid,
			Available: item.Available,
			Note:      item.Note,
			Middle:    item.Middle,
			Last:      item.Last,
			AcademicYear: &response.AcademicYear{
				Uuid: item.AcademicYear.Uuid,
				Name: yearName,
			},
			Subject: &response.Subject{
				Uuid: item.Subject.Uuid,
				Name: item.Subject.Name,
				Code: item.Subject.Code,
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *AcademicPlanHandler) UpdateAcademicPlan(c *gin.Context) {
	var body request.UpdateAcademicPlan
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

	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	uuid := c.Param("uuid")
	if err := h.Service.UpdateAcademicPlan(userUuid, uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response.SUCCESS_RES("RPS Berhasil Diperbarui"))
}

func (h *AcademicPlanHandler) DeleteAcademicPlan(c *gin.Context) {
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	uuid := c.Param("uuid")
	if err := h.Service.DeleteAcademicPlan(userUuid, uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("RPS Berhasil Dihapus"))
}

func (h *AcademicPlanHandler) UpdateMiddle(c *gin.Context) {
	var body request.AcademicPlanMiddle
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

	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	uuid := c.Param("uuid")
	if err := h.Service.UpdateOne(userUuid, uuid, "middle", body.Middle); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Kesesuaian Berhasil Diperbarui"))
}

func (h *AcademicPlanHandler) UpdateLast(c *gin.Context) {
	var body request.AcademicPlanLast
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

	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	uuid := c.Param("uuid")
	if err := h.Service.UpdateOne(userUuid, uuid, "last", body.Last); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Kesesuaian Berhasil Diperbarui"))
}
