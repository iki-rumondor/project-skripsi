package request

type Prodi struct {
	Nama      string `json:"nama" valid:"required"`
	Kaprodi   string `json:"kaprodi" valid:"required"`
	JurusanID uint   `json:"jurusan" valid:"required"`
}
