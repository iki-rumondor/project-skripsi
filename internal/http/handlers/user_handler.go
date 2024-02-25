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

func (h *UserHandler) GetDashboardAdmin(c *gin.Context) {
	res, err := h.Service.GetDashboardAdmin()
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

	res := []response.MonevAmount{
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

func (h *UserHandler) GetSettings(c *gin.Context) {

	result, err := h.Service.GetAll("settings")
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var resp []response.Setting
	for _, item := range result {
		resp = append(resp, response.Setting{
			ID:    item["id"],
			Name:  item["name"],
			Value: item["value"],
		})
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *UserHandler) GetDepartmentMonev(c *gin.Context) {
	departmentUuid := c.Param("departmentUuid")
	yearUuid := c.Param("yearUuid")

	result, err := h.Service.CountDepartmentMonev(departmentUuid, yearUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	resp := map[string]interface{}{
		"first_monev": []int{
			result["plans"],
			result["modules"],
			result["tools"],
			result["skills"],
			result["facilities"],
		},
		"middle_monev": []int{
			result["t_att"],
			result["s_att"],
			result["mid_plans"],
		},
		"middle_last_monev": []int{
			result["lt_att"],
			result["ls_att"],
			result["last_plans"],
		},
		"last_monev": []int{
			result["final"],
			result["passed"],
			result["grade"],
		},
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *UserHandler) UpdateStepMonev(c *gin.Context) {
	var body request.StepMonev
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if err := h.Service.Update(body.ID, "settings", "value", body.Step); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Tahapan Monev Berhasil Diperbarui"))
}

func (h *UserHandler) GetDepartmentData(c *gin.Context) {
	userUuid := c.Param("userUuid")

	user, err := h.Service.GetUser("uuid", userUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	resp := &response.Department{
		Uuid:      user.Department.Uuid,
		Name:      user.Department.Name,
		Head:      user.Department.Head,
		CreatedAt: user.Department.CreatedAt,
		UpdatedAt: user.Department.UpdatedAt,
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}
