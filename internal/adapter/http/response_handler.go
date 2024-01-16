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

type ResponseHandler struct {
	Service *application.ResponseService
}

func NewResponseHandler(service *application.ResponseService) *ResponseHandler {
	return &ResponseHandler{
		Service: service,
	}
}

func (h *ResponseHandler) CreateResponse(c *gin.Context) {
	var body request.AssResponse
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

	user, err := h.Service.Repo.UserRepo.FindUserByUuid(body.UserUuid)
	if err != nil {
		utils.HandleError(c, &response.Error{
			Code:    404,
			Message: "Gagal Mendapatkan User",
		})
		return
	}

	question, err := h.Service.Repo.AssQuestionRepo.FindAssQuestionByUuid(body.QuestionUuid)
	if err != nil {
		utils.HandleError(c, &response.Error{
			Code:    404,
			Message: "Gagal Mendapatkan Question",
		})
		return
	}

	model := domain.Response{
		Uuid:       uuid.NewString(),
		UserID:     user.ID,
		QuestionID: question.ID,
		Response:   body.Response,
	}

	if err := h.Service.CreateResponse(&model); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.Message{
		Success: true,
		Message: "Respon berhasil ditambahkan",
	})
}

func (h *ResponseHandler) GetAllResponse(c *gin.Context) {

	urlPath := c.Request.URL.Path

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))

	pagination := domain.Pagination{
		Limit: limit,
		Page:  page,
	}

	res, err := h.Service.GetAllResponse(urlPath, &pagination)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, &response.SuccessResponse{
		Success: true,
		Data:    res,
	})
}

func (h *ResponseHandler) GetResponseByUuid(c *gin.Context) {

	uuid := c.Param("uuid")

	result, err := h.Service.GetResponse(uuid)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var res = response.AssResponse{
		Uuid:     result.Uuid,
		Response: result.Response,
		User: &response.User{
			Uuid:     result.User.Uuid,
			Username: result.User.Username,
			Role:     result.User.Role,
		},
		Question: &response.AssQuestion{
			Uuid:     result.Question.Uuid,
			Question: result.Question.Question,
			AssessmentType: &response.AssType{
				Uuid:      result.Question.AssessmentType.Uuid,
				Type:      result.Question.AssessmentType.Type,
				CreatedAt: result.Question.AssessmentType.CreatedAt,
				UpdatedAt: result.Question.AssessmentType.UpdatedAt,
			},
			CreatedAt: result.Question.CreatedAt,
			UpdatedAt: result.Question.UpdatedAt,
		},
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	c.JSON(http.StatusOK, &response.SuccessResponse{
		Success: true,
		Data:    res,
	})
}

func (h *ResponseHandler) UpdateResponse(c *gin.Context) {

	uuid := c.Param("uuid")

	result, err := h.Service.GetResponse(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var body request.AssResponse
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

	user, err := h.Service.Repo.UserRepo.FindUserByUuid(body.UserUuid)
	if err != nil {
		utils.HandleError(c, &response.Error{
			Code:    404,
			Message: "Gagal Mendapatkan User",
		})
		return
	}

	question, err := h.Service.Repo.AssQuestionRepo.FindAssQuestionByUuid(body.QuestionUuid)
	if err != nil {
		utils.HandleError(c, &response.Error{
			Code:    404,
			Message: "Gagal Mendapatkan Question",
		})
		return
	}

	model := domain.Response{
		ID:         result.ID,
		UserID:     user.ID,
		QuestionID: question.ID,
		Response:   body.Response,
	}

	if err := h.Service.UpdateResponse(&model); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.Message{
		Success: true,
		Message: "Respon berhasil diperbarui",
	})
}

func (h *ResponseHandler) DeleteResponse(c *gin.Context) {

	uuid := c.Param("uuid")

	result, err := h.Service.GetResponse(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	model := domain.Response{
		ID: result.ID,
	}

	if err := h.Service.DeleteResponse(&model); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.Message{
		Success: true,
		Message: "Respon berhasil dihapus",
	})
}
