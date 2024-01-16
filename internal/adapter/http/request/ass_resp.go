package request

type AssResponse struct {
	UserUuid     string `json:"user_uuid" valid:"required~field user_uuid is required"`
	QuestionUuid string `json:"question_uuid" valid:"required~field question_uuid is required"`
	Response     bool   `json:"response" valid:"required~field response is required"`
}
