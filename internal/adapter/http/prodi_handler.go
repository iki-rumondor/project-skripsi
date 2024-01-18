package customHTTP

import (
	"errors"
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
	"gorm.io/gorm"
)

type ProdiHandler struct {
	Service *application.ProdiService
}

func NewProdiHandler(service *application.ProdiService) *ProdiHandler {
	return &ProdiHandler{
		Service: service,
	}
}

func (h *ProdiHandler) GetAllProdi(c *gin.Context) {

	prodi, err := h.Service.GetAllProdi()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	var res = []*response.Prodi{}

	for _, p := range *prodi {
		res = append(res, &response.Prodi{
			Uuid:       p.Uuid,
			Nama:       p.Nama,
			Kaprodi:    p.Kaprodi,
			Credential: p.Credential,
			Jurusan: &response.JurusanData{
				ID:        p.Jurusan.ID,
				Nama:      p.Jurusan.Nama,
				CreatedAt: p.CreatedAt,
			},
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    res,
	})
}

func (h *ProdiHandler) GetProdiByUuid(c *gin.Context) {

	uuid := c.Param("uuid")

	prodi, err := h.Service.GetProdiByUuid(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var subjects []response.Subject

	for _, item := range *prodi.Subject {
		subjects = append(subjects, response.Subject{
			Uuid:      item.Uuid,
			Name:      item.Name,
			Code:      item.Code,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	var res = response.Prodi{
		Uuid:       prodi.Uuid,
		Nama:       prodi.Nama,
		Kaprodi:    prodi.Kaprodi,
		Credential: prodi.Credential,
		Jurusan: &response.JurusanData{
			ID:        prodi.Jurusan.ID,
			Nama:      prodi.Jurusan.Nama,
			CreatedAt: prodi.CreatedAt,
		},
		Subject:   &subjects,
		CreatedAt: prodi.CreatedAt,
		UpdatedAt: prodi.UpdatedAt,
	}

	c.JSON(http.StatusOK, res)
}

func (h *ProdiHandler) CreateProdi(c *gin.Context) {
	var body request.Prodi
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	jurusanID, err := strconv.Atoi(body.JurusanID)
	if err != nil {
		utils.HandleError(c, &response.Error{
			Code:    404,
			Message: "Data jurusan tidak valid",
		})
	}

	prodi := domain.Prodi{
		Uuid:       uuid.NewString(),
		Credential: utils.GenerateRandomString(15),
		Nama:       body.Nama,
		Kaprodi:    body.Kaprodi,
		JurusanID:  uint(jurusanID),
	}

	if err := h.Service.CreateProdi(&prodi); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, response.Message{
				Success: false,
				Message: err.Error(),
			})
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Message{
		Success: true,
		Message: "Prodi has been created successfully",
	})
}

func (h *ProdiHandler) UpdateProdi(c *gin.Context) {

	var body request.Prodi
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	uuid := c.Param("uuid")

	if err := h.Service.UpdateProdi(uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.Message{
		Success: true,
		Message: "Prodi has been updated successfully",
	})
}

func (h *ProdiHandler) DeleteProdi(c *gin.Context) {

	uuid := c.Param("uuid")

	if err := h.Service.DeleteProdi(uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.Message{
		Success: true,
		Message: "Prodi has been deleted successfully",
	})
}
