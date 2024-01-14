package response

import "time"

type AssType struct {
	Uuid      string    `json:"uuid"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
