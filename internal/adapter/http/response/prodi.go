package response

import "time"

type Prodi struct {
	Uuid       string       `json:"uuid"`
	Nama       string       `json:"nama"`
	Kaprodi    string       `json:"kaprodi"`
	Credential string       `json:"credential"`
	Jurusan    *JurusanData `json:"jurusan"`
	Subject    *[]Subject   `json:"subjects"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
}

type Subject struct {
	Uuid      string    `json:"uuid"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	RPS       *bool     `json:"rps"`
	Prodi     *Prodi    `json:"prodi"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
