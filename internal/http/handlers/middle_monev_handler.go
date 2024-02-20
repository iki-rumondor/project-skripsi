package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/http/response"
	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/models"
	"github.com/iki-rumondor/go-monev/internal/utils"
)

type MiddleMonevHandler struct {
	Service interfaces.MiddleMonevServiceInterface
}

func NewMiddleMonevHandler(service interfaces.MiddleMonevServiceInterface) interfaces.MiddleMonevHandlerInterface {
	return &MiddleMonevHandler{
		Service: service,
	}
}

func (h *MiddleMonevHandler) CountTotalMonev(c *gin.Context) {
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	yearUuid := c.Param("yearUuid")

	result, err := h.Service.CountTotalMonev(userUuid, yearUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	resp := []response.MonevAmount{
		{
			Name:   "Persentase Kehadiran Dosen Dalam Mengajar",
			Amount: result["t_att"],
		},
		{
			Name:   "Persentase Kehadiran Mahasiswa",
			Amount: result["s_att"],
		},
		{
			Name:   "Kesesuaian Mengajar Dosen dengan RPS",
			Amount: result["plans"],
		},
	}

	c.JSON(http.StatusCreated, response.DATA_RES(resp))
}

func (h *MiddleMonevHandler) CreateTeacherAttendence(c *gin.Context) {
	var body request.CreateTeacherAttendence
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

	if err := h.Service.CreateTeacherAttendence(userUuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Kehadiran Dosen Berhasil Ditambahkan"))
}

func (h *MiddleMonevHandler) GetTeacherAttendences(c *gin.Context) {

	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	yearUuid := c.Param("yearUuid")

	result, err := h.Service.GetTeacherAttendences(userUuid, yearUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var resp []*response.TeacherAttendence
	for _, item := range *result {
		resp = append(resp, &response.TeacherAttendence{
			Uuid:        item.Uuid,
			Middle:      fmt.Sprintf("%d", item.Middle),
			Last:        fmt.Sprintf("%d", item.Last),
			GeadeOnTime: item.GradeOnTime,
			AcademicYear: &response.AcademicYear{
				Uuid: item.AcademicYear.Uuid,
				Name: item.AcademicYear.Name,
			},
			Subject: &response.Subject{
				Uuid: item.Subject.Uuid,
				Name: item.Subject.Name,
				Code: item.Subject.Code,
			},
			Teacher: &response.Teacher{
				Uuid: item.Teacher.Uuid,
				Name: item.Teacher.Name,
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	c.JSON(http.StatusCreated, response.DATA_RES(resp))
}

func (h *MiddleMonevHandler) DeleteTeacherAttendence(c *gin.Context) {

	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	uuid := c.Param("uuid")

	if err := h.Service.DeleteTeacherAttendence(userUuid, uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Kehadiran Dosen Berhasil Dihapus"))
}

func (h *MiddleMonevHandler) CreateStudentAttendence(c *gin.Context) {
	var body request.CreateStudentAttendence
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

	if err := h.Service.CreateStudentAttendence(userUuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Kehadiran Mahasiswa Berhasil Ditambahkan"))
}

func (h *MiddleMonevHandler) GetStudentAttendences(c *gin.Context) {
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	yearUuid := c.Param("yearUuid")

	result, err := h.Service.GetStudentAttendences(userUuid, yearUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var resp []*response.StudentAttendence
	for _, item := range *result {
		resp = append(resp, &response.StudentAttendence{
			Uuid:          item.Uuid,
			StudentAmount: fmt.Sprintf("%d", item.StudentAmount),
			Middle:        fmt.Sprintf("%d", item.Middle),
			Last:          fmt.Sprintf("%d", item.Last),
			PassedAmount:  fmt.Sprintf("%d", item.PassedAmount),
			FinalAmount:   fmt.Sprintf("%d", item.FinalAmount),
			AcademicYear: &response.AcademicYear{
				Uuid: item.AcademicYear.Uuid,
				Name: item.AcademicYear.Name,
			},
			Subject: &response.Subject{
				Uuid: item.Subject.Uuid,
				Name: item.Subject.Name,
				Code: item.Subject.Code,
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	c.JSON(http.StatusCreated, response.DATA_RES(resp))
}

func (h *MiddleMonevHandler) DeleteStudentAttendence(c *gin.Context) {
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	uuid := c.Param("uuid")

	if err := h.Service.DeleteStudentAttendence(userUuid, uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Kehadiran Mahasiswa Berhasil Dihapus"))
}

func (h *MiddleMonevHandler) UpdateLastTeacherAttendence(c *gin.Context) {
	var body request.LastTeacherAttendence
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

	uuid := c.Param("uuid")

	last, _ := strconv.Atoi(body.Last)

	model := models.TeacherAttendence{
		Last: uint(last),
	}

	if err := h.Service.UpdateTeacherAttendence(userUuid, uuid, &model); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Kehadiran Dosen Berhasil Diperbarui"))
}

func (h *MiddleMonevHandler) UpdateLastStudentAttendence(c *gin.Context) {
	var body request.LastStudentAttendence
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

	uuid := c.Param("uuid")

	last, _ := strconv.Atoi(body.Last)

	model := models.StudentAttendence{
		Last: uint(last),
	}

	if err := h.Service.UpdateStudentAttendence(userUuid, uuid, &model); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Kehadiran Mahasiswa Berhasil Diperbarui"))
}

func (h *MiddleMonevHandler) GetTeacherAttendence(c *gin.Context) {
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	uuid := c.Param("uuid")

	result, err := h.Service.GetTeacherAttendence(userUuid, uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var resp = response.TeacherAttendence{
		Uuid:   result.Uuid,
		Middle: fmt.Sprintf("%d", result.Middle),
		Last:   fmt.Sprintf("%d", result.Last),
		AcademicYear: &response.AcademicYear{
			Uuid: result.AcademicYear.Uuid,
			Name: result.AcademicYear.Name,
		},
		Subject: &response.Subject{
			Uuid: result.Subject.Uuid,
			Name: result.Subject.Name,
			Code: result.Subject.Code,
		},
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	c.JSON(http.StatusCreated, response.DATA_RES(resp))
}
