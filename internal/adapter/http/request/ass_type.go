package request

type AssType struct{
	Type string `json:"type" valid:"required~field type is required"`
}