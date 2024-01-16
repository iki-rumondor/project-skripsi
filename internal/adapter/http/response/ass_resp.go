package response

import "time"

type AssResponse struct {
	Uuid      string       `json:"uuid"`
	Response  bool         `json:"response"`
	User      *User        `json:"user"`
	Question  *AssQuestion `json:"question"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}
