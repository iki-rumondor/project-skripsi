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

type TeacherSkillHandler struct {
	Service interfaces.TeacherSkillServiceInterface
}

func NewTeacherSkillHandler(service interfaces.TeacherSkillServiceInterface) interfaces.TeacherSkillHandlerInterface {
	return &TeacherSkillHandler{
		Service: service,
	}
}

func (h *TeacherSkillHandler) CreateTeacherSkill(c *gin.Context) {
	var body request.TeacherSkill
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

	if err := h.Service.CreateTeacherSkill(userUuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Kemampuan Dosen Berhasil Ditambahkan"))
}

func (h *TeacherSkillHandler) GetAllTeacherSkills(c *gin.Context) {
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	result, err := h.Service.GetAllTeacherSkills(userUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var resp []*response.TeacherSkill
	for _, item := range *result {
		resp = append(resp, &response.TeacherSkill{
			Uuid:  item.Uuid,
			Skill: item.Skill,
			Teacher: &response.Teacher{
				Uuid:      item.Teacher.Uuid,
				Name:      item.Teacher.Name,
				CreatedAt: item.Teacher.CreatedAt,
				UpdatedAt: item.Teacher.UpdatedAt,
			},
			Subject: &response.Subject{
				Uuid:      item.Subject.Uuid,
				Name:      item.Subject.Name,
				Code:      item.Subject.Code,
				Practical: item.Subject.Practical,
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *TeacherSkillHandler) GetByYear(c *gin.Context) {
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	yearUuid := c.Param("yearUuid")

	result, err := h.Service.GetTeacherSkillsByYear(userUuid, yearUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var resp []*response.TeacherSkill
	for _, item := range *result {
		resp = append(resp, &response.TeacherSkill{
			Uuid:  item.Uuid,
			Skill: item.Skill,
			Teacher: &response.Teacher{
				Uuid:      item.Teacher.Uuid,
				Name:      item.Teacher.Name,
				CreatedAt: item.Teacher.CreatedAt,
				UpdatedAt: item.Teacher.UpdatedAt,
			},
			Subject: &response.Subject{
				Uuid:      item.Subject.Uuid,
				Name:      item.Subject.Name,
				Code:      item.Subject.Code,
				Practical: item.Subject.Practical,
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *TeacherSkillHandler) GetTeacherSkill(c *gin.Context) {
	uuid := c.Param("uuid")
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}
	result, err := h.Service.GetTeacherSkill(userUuid, uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	resp := &response.TeacherSkill{
		Uuid:  result.Uuid,
		Skill: result.Skill,
		Teacher: &response.Teacher{
			Uuid:      result.Teacher.Uuid,
			Name:      result.Teacher.Name,
			CreatedAt: result.Teacher.CreatedAt,
			UpdatedAt: result.Teacher.UpdatedAt,
		},
		Subject: &response.Subject{
			Uuid:      result.Subject.Uuid,
			Name:      result.Subject.Name,
			Code:      result.Subject.Code,
			Practical: result.Subject.Practical,
		},
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *TeacherSkillHandler) UpdateTeacherSkill(c *gin.Context) {
	var body request.UpdateTeacherSkill
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: err.Error(),
		})
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
	if err := h.Service.UpdateTeacherSkill(userUuid, uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response.SUCCESS_RES("Kemampuan Dosen Berhasil Diperbarui"))
}

func (h *TeacherSkillHandler) DeleteTeacherSkill(c *gin.Context) {
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	uuid := c.Param("uuid")
	if err := h.Service.DeleteTeacherSkill(userUuid, uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Kemampuan Dosen Berhasil Dihapus"))
}
