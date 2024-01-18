package customHTTP

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
)

func (h *ProdiHandler) CreateSubject(c *gin.Context) {
	var body request.Subject
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, badRequestError)
		return
	}

	userUuid := c.GetString("user_uuid")

	if err := h.Service.CreateSubject(userUuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.Message{
		Success: true,
		Message: "Mata Kuliah Berhasil Ditambahkan",
	})
}

func (h *ProdiHandler) GetAllSubjects(c *gin.Context) {

	result, err := h.Service.GetAllSubject()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var res = []*response.Subject{}

	for _, item := range *result {
		res = append(res, &response.Subject{
			Uuid: item.Uuid,
			Name: item.Name,
			Code: item.Code,
			Prodi: &response.Prodi{
				Uuid:      item.Prodi.Uuid,
				Nama:      item.Prodi.Nama,
				Kaprodi:   item.Prodi.Kaprodi,
				CreatedAt: item.Prodi.CreatedAt,
				UpdatedAt: item.Prodi.UpdatedAt,
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Data:    res,
	})
}

func (h *ProdiHandler) GetProdiSubjects(c *gin.Context) {

	uuid := c.Param("uuid")

	result, err := h.Service.GetProdiSubjects(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var res = []*response.Subject{}

	for _, item := range *result {
		res = append(res, &response.Subject{
			Uuid:      item.Uuid,
			Name:      item.Name,
			Code:      item.Code,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Data:    res,
	})
}

func (h *ProdiHandler) GetSubjectByUuid(c *gin.Context) {

	uuid := c.Param("uuid")

	result, err := h.Service.GetSubjectByUuid(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var res = response.Subject{
		Uuid: result.Uuid,
		Name: result.Name,
		Code: result.Code,
		Prodi: &response.Prodi{
			Uuid:      result.Prodi.Uuid,
			Nama:      result.Prodi.Nama,
			Kaprodi:   result.Prodi.Kaprodi,
			CreatedAt: result.Prodi.CreatedAt,
			UpdatedAt: result.Prodi.UpdatedAt,
		},
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Data:    res,
	})
}

func (h *ProdiHandler) UpdateSubject(c *gin.Context) {

	var body request.Subject
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, badRequestError)
		return
	}

	uuid := c.Param("uuid")
	userUuid := c.GetString("user_uuid")

	if err := h.Service.UpdateSubject(userUuid, uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.Message{
		Success: true,
		Message: "Mata Kuliah Berhasil Diupdate",
	})
}

func (h *ProdiHandler) DeleteSubject(c *gin.Context) {

	uuid := c.Param("uuid")
	userUuid := c.GetString("user_uuid")

	if err := h.Service.DeleteSubject(userUuid, uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.Message{
		Success: true,
		Message: "Mata Kuliah Berhasil Dihapus",
	})
}
