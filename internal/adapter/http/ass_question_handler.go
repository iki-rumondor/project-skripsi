package customHTTP

import (
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
)

type AssQuestionHandler struct {
	Service *application.AssQuestionService
}

func NewAssQuestionHandler(service *application.AssQuestionService) *AssQuestionHandler {
	return &AssQuestionHandler{
		Service: service,
	}
}

func (h *AssQuestionHandler) CreateAssQuestion(c *gin.Context) {
	var body request.AssQuestion
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

	result, err := h.Service.Repo.AssTypeRepo.FindAssTypeByUuid(body.TypeUuid)
	if err != nil {
		utils.HandleError(c, &response.Error{
			Code:    404,
			Message: "Gagal Mendapatkan Tipe Penilaian",
		})
		return
	}

	model := domain.AssessmentQuestion{
		Uuid:             uuid.NewString(),
		Question:         body.Question,
		AssessmentTypeID: result.ID,
	}

	if err := h.Service.CreateAssQuestion(&model); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.Message{
		Success: true,
		Message: "Pertanyaan berhasil ditambahkan",
	})
}

func (h *AssQuestionHandler) GetAllAssQuestion(c *gin.Context) {

	urlPath := c.Request.URL.Path

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))

	pagination := domain.Pagination{
		Limit: limit,
		Page:  page,
	}

	res, err := h.Service.GetAllAssQuestion(urlPath, &pagination)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, &response.SuccessResponse{
		Success: true,
		Data:    res,
	})
}

func (h *AssQuestionHandler) GetAssQuestionByUuid(c *gin.Context) {

	uuid := c.Param("uuid")

	result, err := h.Service.GetAssQuestion(uuid)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var res = response.AssQuestion{
		Uuid:     result.Uuid,
		Question: result.Question,
		AssessmentType: &response.AssType{
			Uuid:      result.AssessmentType.Uuid,
			Type:      result.AssessmentType.Type,
			CreatedAt: result.AssessmentType.CreatedAt,
			UpdatedAt: result.AssessmentType.UpdatedAt,
		},
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	c.JSON(http.StatusOK, &response.SuccessResponse{
		Success: true,
		Data:    res,
	})
}

func (h *AssQuestionHandler) UpdateAssQuestion(c *gin.Context) {

	uuid := c.Param("uuid")

	result, err := h.Service.GetAssQuestion(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var body request.AssQuestion
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

	res, err := h.Service.Repo.AssTypeRepo.FindAssTypeByUuid(body.TypeUuid)
	if err != nil {
		utils.HandleError(c, &response.Error{
			Code:    404,
			Message: "Gagal Mendapatkan Tipe Penilaian",
		})
		return
	}

	model := domain.AssessmentQuestion{
		ID:               result.ID,
		Question:         body.Question,
		AssessmentTypeID: res.ID,
	}

	if err := h.Service.UpdateAssQuestion(&model); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.Message{
		Success: true,
		Message: "Pertanyaan berhasil diperbarui",
	})
}

func (h *AssQuestionHandler) DeleteAssQuestion(c *gin.Context) {

	uuid := c.Param("uuid")

	result, err := h.Service.GetAssQuestion(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	model := domain.AssessmentQuestion{
		ID: result.ID,
	}

	if err := h.Service.DeleteAssQuestion(&model); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.Message{
		Success: true,
		Message: "Pertanyaan berhasil dihapus",
	})
}
