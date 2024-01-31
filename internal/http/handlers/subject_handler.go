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

type SubjectHandler struct {
	Service interfaces.SubjectServiceInterface
}

func NewSubjectHandler(service interfaces.SubjectServiceInterface) interfaces.SubjectHandlerInterface {
	return &SubjectHandler{
		Service: service,
	}
}

func (h *SubjectHandler) CreateSubject(c *gin.Context) {
	var body request.Subject
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
	
	body.UserUuid = userUuid

	if err := h.Service.CreateSubject(&body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Mata Kuliah Berhasil Ditambahkan"))
}

func (h *SubjectHandler) GetAllSubjects(c *gin.Context) {
	subjects, err := h.Service.GetAllSubjects()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var resp []*response.Subject
	for _, item := range *subjects {
		resp = append(resp, &response.Subject{
			Uuid: item.Uuid,
			Name: item.Name,
			Code: item.Code,
			Department: &response.Department{
				Uuid: item.Department.Uuid,
				Name: item.Department.Name,
				Head: item.Department.Head,
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *SubjectHandler) GetSubject(c *gin.Context) {
	uuid := c.Param("uuid")
	subject, err := h.Service.GetSubject(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	resp := &response.Subject{
		Uuid: subject.Uuid,
		Name: subject.Name,
		Code: subject.Code,
		Department: &response.Department{
			Uuid: subject.Department.Uuid,
			Name: subject.Department.Name,
			Head: subject.Department.Head,
		},
		CreatedAt: subject.CreatedAt,
		UpdatedAt: subject.UpdatedAt,
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *SubjectHandler) UpdateSubject(c *gin.Context) {
	var body request.Subject
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
	body.UserUuid = userUuid
	
	uuid := c.Param("uuid")
	if err := h.Service.UpdateSubject(uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Mata Kuliah Berhasil Diperbarui"))
}

func (h *SubjectHandler) DeleteSubject(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := h.Service.DeleteSubject(uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Mata Kuliah Berhasil Dihapus"))
}
