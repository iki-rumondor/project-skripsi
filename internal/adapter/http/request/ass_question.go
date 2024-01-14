package request

type AssQuestion struct {
	Question string `json:"question" valid:"required~field question is required"`
	TypeUuid string `json:"type_uuid" valid:"required~field type_uuid is required"`
}
