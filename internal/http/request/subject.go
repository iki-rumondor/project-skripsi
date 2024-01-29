package request

type Subject struct {
	DepartmentUuid string `json:"department_id" binding:"required"`
	Name           string `json:"name" binding:"required"`
	Code           string `json:"code" binding:"required"`
}
