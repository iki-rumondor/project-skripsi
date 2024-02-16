package handlers

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/http/response"
	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/utils"
)

type UserHandler struct {
	Service interfaces.UserServiceInterface
}

func NewUserHandler(service interfaces.UserServiceInterface) interfaces.UserHandlerInterface {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) SignIn(c *gin.Context) {
	var body request.SignIn
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	jwt, err := h.Service.VerifyUser(&body)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": jwt,
	})
}

func (h *UserHandler) GetCountSubjects(c *gin.Context) {
	res, err := h.Service.GetCountSubjects()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(res))
}

func (h *UserHandler) CountMonevByYear(c *gin.Context) {

	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	yearUuid := c.Param("yearUuid")

	result, err := h.Service.CountMonevByYear(userUuid, yearUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	res := []response.MonevCount{
		{
			Name:   "Ketersediaan RPS",
			Amount: result["plans"],
		},
		{
			Name:   "Ketersediaan Modul Praktikum",
			Amount: result["modules"],
		},
		{
			Name:   "Ketersediaan Alat Praktikum",
			Amount: result["tools"],
		},
		{
			Name:   "Kesesuaian Kemampuan Dosen Dengan Mata Kuliah",
			Amount: result["skills"],
		},
		{
			Name:   "Fasilitas, Sarana, dan Prasarana",
			Amount: result["facilities"],
		},
	}

	c.JSON(http.StatusOK, response.DATA_RES(res))
}
