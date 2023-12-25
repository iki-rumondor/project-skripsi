package response

import "time"

type Prodi struct {
	ID        uint      `json:"id"`
	Nama      string    `json:"nama"`
	Kaprodi   string    `json:"kaprodi"`
	JurusanID   uint    `json:"jurusan_id"`
	Jurusan   string    `json:"jurusan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}