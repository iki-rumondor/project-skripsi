package customHTTP

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
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
			ID:        p.ID,
			Nama:      p.Nama,
			Kaprodi:   p.Kaprodi,
			Jurusan:   p.Jurusan.Nama,
			JurusanID: p.Jurusan.ID,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    res,
	})
}

func (h *ProdiHandler) GetProdiByID(c *gin.Context) {

	urlParam := c.Param("id")
	prodiID, err := strconv.Atoi(urlParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: "please check the url and ensure it follows the correct format",
		})
		return
	}

	prodi, err := h.Service.GetProdiByID(uint(prodiID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	var res = response.Prodi{
		ID:        prodi.ID,
		Nama:      prodi.Nama,
		Kaprodi:   prodi.Kaprodi,
		Jurusan:   prodi.Jurusan.Nama,
		JurusanID: prodi.Jurusan.ID,
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
		Nama:      body.Nama,
		Kaprodi:   body.Kaprodi,
		JurusanID: uint(jurusanID),
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

	urlParam := c.Param("id")
	prodiID, err := strconv.Atoi(urlParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: "please check the url and ensure it follows the correct format",
		})
		return
	}

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
		ID:        uint(prodiID),
		Nama:      body.Nama,
		Kaprodi:   body.Kaprodi,
		JurusanID: uint(jurusanID),
	}

	if err := h.Service.UpdateProdi(&prodi); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, response.Message{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Message{
		Success: true,
		Message: "Prodi has been updated successfully",
	})
}

func (h *ProdiHandler) DeleteProdi(c *gin.Context) {

	urlParam := c.Param("id")
	prodiID, err := strconv.Atoi(urlParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: "please check the url and ensure it follows the correct format",
		})
		return
	}

	prodi := domain.Prodi{
		ID: uint(prodiID),
	}

	if err := h.Service.DeleteProdi(&prodi); err != nil {
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
		Message: "Prodi has been deleted successfully",
	})
}
