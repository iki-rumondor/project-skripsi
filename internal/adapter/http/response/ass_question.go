package response

import "time"

type AssQuestion struct {
	Uuid           string    `json:"uuid"`
	Question       string    `json:"question"`
	AssessmentType *AssType  `json:"type"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
