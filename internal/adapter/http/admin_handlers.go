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
	"gorm.io/gorm"
)

type AdminHandler struct {
	Service *application.AdminService
}

func NewAdminHandler(service *application.AdminService) *AdminHandler {
	return &AdminHandler{
		Service: service,
	}
}

func (h *AdminHandler) GetAllIndikator(c *gin.Context) {

	result, err := h.Service.GetAllIndikator()
	
	if err != nil {
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

	var res = []*response.Indikator{}

	
	for _,item := range *result{
		res = append(res, &response.Indikator{
			ID: item.ID,
			Description: item.Description,
			TypeID: item.IndikatorType.ID,
			Type: item.IndikatorType.Description,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &res,
	})
}

func (h *AdminHandler) GetIndikator(c *gin.Context) {

	urlParam := c.Param("id")
	id, err := strconv.Atoi(urlParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Message: "please check the url and ensure it follows the correct format",
		})
		return
	}

	result, err := h.Service.GetIndikator(uint(id))
	
	if err != nil {
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

	var res = response.Indikator{
		ID: result.ID,
		Description: result.Description,
		TypeID: result.IndikatorType.ID,
		Type: result.IndikatorType.Description,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	
	c.JSON(http.StatusOK, res)
}

func (h *AdminHandler) CreateIndikator(c *gin.Context) {

	var body request.Indikator
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Message: err.Error(),
		})
		return
	}

	if err := h.Service.CreateIndikator(&body); err != nil {
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Indikator has been created successfully",
	})
}

func (h *AdminHandler) UpdateIndikator(c *gin.Context) {

	var body request.Indikator
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Message: err.Error(),
		})
		return
	}

	urlParam := c.Param("id")
	id, err := strconv.Atoi(urlParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Message: "please check the url and ensure it follows the correct format",
		})
		return
	}

	if err := h.Service.UpdateIndikator(uint(id), &body); err != nil {
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Indikator has been updated successfully",
	})
}

func (h *AdminHandler) DeleteIndikator(c *gin.Context) {

	urlParam := c.Param("id")
	id, err := strconv.Atoi(urlParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Message: "please check the url and ensure it follows the correct format",
		})
		return
	}

	if err := h.Service.DeleteIndikator(uint(id)); err != nil {
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Indikator has been deleted successfully",
	})
}

func (h *AdminHandler) CreateInstrumenType(c *gin.Context) {

	var body request.CreateInstrumenType
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Message: err.Error(),
		})
		return
	}

	if err := h.Service.CreateInstrumenType(&body); err != nil {
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Your instrumen type has been created successfully",
	})
}

func (h *AdminHandler) GetAllInstrumenType(c *gin.Context) {

	result, err := h.Service.GetAllInstrumenType()
	
	if err != nil {
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

	var res = []*response.InstrumenType{}

	for _,item := range *result{
		res = append(res, &response.InstrumenType{
			ID: item.ID,
			Name: item.Name,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &res,
	})
}

func (h *AdminHandler) DeleteInstrumenType(c *gin.Context) {

	urlParam := c.Param("id")
	id, err := strconv.Atoi(urlParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Message: "please check the url and ensure it follows the correct format",
		})
		return
	}

	if err := h.Service.DeleteInstrumenType(uint(id)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, response.ErrorMessage{
				Message: err.Error(),
			})
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorMessage{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Instrumen type has been deleted successfully",
	})
}

func (h *AdminHandler) GetAllIndikatorType(c *gin.Context) {

	result, err := h.Service.GetAllIndikatorType()
	
	if err != nil {
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

	var res = []*response.IndikatorType{}

	for _,item := range *result{
		res = append(res, &response.IndikatorType{
			ID: item.ID,
			Description: item.Description,
			Instrument: item.InstrumenType.Name,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &res,
	})
}

func (h *AdminHandler) CreateIndikatorType(c *gin.Context) {

	var body request.CreateIndikatorType
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Message: err.Error(),
		})
		return
	}

	if err := h.Service.CreateIndikatorType(&body); err != nil {
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Your indikator type has been created successfully",
	})
}

func (h *AdminHandler) DeleteIndikatorType(c *gin.Context) {

	urlParam := c.Param("id")
	id, err := strconv.Atoi(urlParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
			Message: "please check the url and ensure it follows the correct format",
		})
		return
	}

	if err := h.Service.DeleteIndikatorType(uint(id)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, response.ErrorMessage{
				Message: err.Error(),
			})
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorMessage{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Indikator type has been deleted successfully",
	})
}
