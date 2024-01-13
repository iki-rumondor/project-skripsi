package customHTTP

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/application"
)

type UtilHandler struct {
	Service *application.UtilService
}

func NewUtilHandler(service *application.UtilService) *UtilHandler {
	return &UtilHandler{
		Service: service,
	}
}

func (h *UtilHandler) GetAllJurusan(c *gin.Context) {

	jurusan, err := h.Service.GetAllJurusan()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	var res = []*response.JurusanData{}

	for _, j := range *jurusan {
		res = append(res, &response.JurusanData{
			ID: j.ID,
			Nama: j.Nama,
			CreatedAt: j.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success" : true,
		"data" : res,
	})
}
