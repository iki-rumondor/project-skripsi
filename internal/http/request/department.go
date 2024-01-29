package request

type Department struct {
	MajorUuid string `json:"major_uuid" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Head      string `json:"head" binding:"required"`
}
