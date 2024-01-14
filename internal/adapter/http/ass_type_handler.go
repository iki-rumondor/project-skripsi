package customHTTP

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
)

type AssTypeHandler struct {
	Service *application.AssTypeService
}

func NewAssTypeHandler(service *application.AssTypeService) *AssTypeHandler {
	return &AssTypeHandler{
		Service: service,
	}
}

func (h *AssTypeHandler) CreateAssType(c *gin.Context) {
	var body request.AssType
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	model := domain.AssessmentType{
		Uuid: uuid.NewString(),
		Type: body.Type,
	}

	if err := h.Service.CreateAssType(&model); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.Message{
		Success: true,
		Message: "Tipe Penilaian berhasil ditambahkan",
	})
}

func (h *AssTypeHandler) GetAllAssType(c *gin.Context) {

	result, err := h.Service.GetAllAssType()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var res = []*response.AssType{}

	for _, item := range *result {
		res = append(res, &response.AssType{
			Uuid:      item.Uuid,
			Type:      item.Type,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, &response.SuccessResponse{
		Success: true,
		Data:    res,
	})
}

func (h *AssTypeHandler) GetAssTypeByUuid(c *gin.Context) {

	uuid := c.Param("uuid")

	result, err := h.Service.GetAssType(uuid)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var res = response.AssType{
		Uuid:      result.Uuid,
		Type:      result.Type,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	c.JSON(http.StatusOK, &response.SuccessResponse{
		Success: true,
		Data:    res,
	})
}

func (h *AssTypeHandler) UpdateAssType(c *gin.Context) {

	uuid := c.Param("uuid")

	result, err := h.Service.GetAssType(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var body request.AssType
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	model := domain.AssessmentType{
		ID:   result.ID,
		Type: body.Type,
	}

	if err := h.Service.UpdateAssType(&model); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.Message{
		Success: true,
		Message: "Tipe penilaian berhasil diperbarui",
	})
}

func (h *AssTypeHandler) DeleteAssType(c *gin.Context) {

	uuid := c.Param("uuid")

	result, err := h.Service.GetAssType(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	AssType := domain.AssessmentType{
		ID: result.ID,
	}

	if err := h.Service.DeleteAssType(&AssType); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.Message{
		Success: true,
		Message: "Tipe penilaian berhasil dihapus",
	})
}
