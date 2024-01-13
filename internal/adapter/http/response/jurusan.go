package response

import "time"

type JurusanData struct {
	ID        uint      `json:"id"`
	Nama      string    `json:"nama"`
	CreatedAt time.Time `json:"created_at"`
}
