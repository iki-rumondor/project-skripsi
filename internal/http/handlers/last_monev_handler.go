package handlers

import (
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/http/response"
	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/utils"
)

type LastMonevHandler struct {
	Service interfaces.LastMonevServiceInterface
}

func NewLastMonevHandler(service interfaces.LastMonevServiceInterface) interfaces.LastMonevHandlerInterface {
	return &LastMonevHandler{
		Service: service,
	}
}

func (h *LastMonevHandler) CountLastMonev(c *gin.Context) {
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	yearUuid := c.Param("yearUuid")

	result, err := h.Service.CountLastMonev(userUuid, yearUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	resp := []response.MonevAmount{
		{
			Name:   "Persentase Kelulusan Mahasiswa Per Mata Kuliah",
			Amount: result["s_att"],
		},
		{
			Name:   "Persentase Keikutsertaan Mahasiswa Mengikuti Ujian Akhir Semester",
			Amount: result["s_att"],
		},
		{
			Name:   "Pemasukan Nilai Akhir Mata Kuliah",
			Amount: result["t_att"],
		},
	}

	c.JSON(http.StatusCreated, response.DATA_RES(resp))
}

func (h *LastMonevHandler) UpdateStudentPass(c *gin.Context) {
	var body request.UpdateStudentPass
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	students, _ := strconv.Atoi(body.StudentAmount)
	passed, _ := strconv.Atoi(body.PassedAmount)

	if passed > students {
		utils.HandleError(c, response.BADREQ_ERR("Jumlah Mahasiswa Lulus Tidak Boleh Lebih Dari Jumlah Total Mahasiswa"))
		return
	}

	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	uuid := c.Param("uuid")

	if err := h.Service.UpdateStudentAttendence(userUuid, uuid, "passed_amount", passed); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Keluluan Mahasiswa Berhasil Ditambahkan"))
}

func (h *LastMonevHandler) UpdateStudentFinal(c *gin.Context) {
	var body request.UpdateStudentFinal
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	students, _ := strconv.Atoi(body.StudentAmount)
	final, _ := strconv.Atoi(body.FinalAmount)

	if final > students {
		utils.HandleError(c, response.BADREQ_ERR("Jumlah Mahasiswa Mengikuti UAS Tidak Boleh Lebih Dari Jumlah Total Mahasiswa"))
		return
	}

	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	uuid := c.Param("uuid")

	if err := h.Service.UpdateStudentAttendence(userUuid, uuid, "final_amount", final); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Keikutsertaan Mahasiswa Berhasil Diperbarui"))
}

func (h *LastMonevHandler) UpdateTeacherGrade(c *gin.Context) {
	var body request.UpdateTeacherGrade
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	uuid := c.Param("uuid")

	if err := h.Service.UpdateTeacherAttendence(userUuid, uuid, "grade_on_time", body.GradeOnTime); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Pemasukan Nilai Berhasil Diperbarui"))
}

