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

type FacilityConditionHandler struct {
	Service interfaces.FacilityConditionServiceInterface
}

func NewFacilityConditionHandler(service interfaces.FacilityConditionServiceInterface) interfaces.FacilityConditionHandlerInterface {
	return &FacilityConditionHandler{
		Service: service,
	}
}

func (h *FacilityConditionHandler) CreateFacilityCondition(c *gin.Context) {
	var body request.FacilityCondition
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

	if err := h.Service.CreateFacilityCondition(userUuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Kondisi Fasilitas Berhasil Ditambahkan"))
}

func (h *FacilityConditionHandler) GetFacilityConditionsByYear(c *gin.Context) {
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	yearUuid := c.Param("yearUuid")

	result, err := h.Service.GetFacilityConditionsByYear(userUuid, yearUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var resp []*response.FacilityCondition
	for _, item := range *result {
		resp = append(resp, &response.FacilityCondition{
			Uuid:     item.Uuid,
			Amount:   fmt.Sprintf("%v", item.Amount),
			Unit:     item.Unit,
			Deactive: fmt.Sprintf("%v", item.Deactive),
			Note:     item.Note,
			AcademicYear: &response.AcademicYear{
				Uuid: item.AcademicYear.Uuid,
				Name: item.AcademicYear.Name,
			},
			Facility: &response.Facility{
				Uuid: item.Facility.Uuid,
				Name: item.Facility.Name,
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FacilityConditionHandler) GetFacilityOptions(c *gin.Context) {
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	yearUuid := c.Param("yearUuid")

	result, err := h.Service.GetFacilityOptions(userUuid, yearUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var resp []*response.Facility
	for _, item := range *result {
		resp = append(resp, &response.Facility{
			Uuid:      item.Uuid,
			Name:      item.Name,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FacilityConditionHandler) GetFacilityCondition(c *gin.Context) {
	uuid := c.Param("uuid")
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}
	result, err := h.Service.GetFacilityCondition(userUuid, uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	resp := &response.FacilityCondition{
		Uuid:     result.Uuid,
		Amount:   fmt.Sprintf("%v", result.Amount),
		Unit:     result.Unit,
		Deactive: fmt.Sprintf("%v", result.Deactive),
		Note:     result.Note,
		AcademicYear: &response.AcademicYear{
			Uuid: result.AcademicYear.Uuid,
			Name: result.AcademicYear.Name,
		},
		Facility: &response.Facility{
			Uuid: result.Facility.Uuid,
			Name: result.Facility.Name,
		},
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *FacilityConditionHandler) UpdateFacilityCondition(c *gin.Context) {
	var body request.UpdateFacilityCondition
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
	if err := h.Service.UpdateFacilityCondition(userUuid, uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response.SUCCESS_RES("Kondisi Fasilitas Berhasil Diperbarui"))
}

func (h *FacilityConditionHandler) DeleteFacilityCondition(c *gin.Context) {
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	uuid := c.Param("uuid")
	if err := h.Service.DeleteFacilityCondition(userUuid, uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Kondisi Fasilitas Berhasil Dihapus"))
}
