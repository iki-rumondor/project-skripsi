package request

type Prodi struct {
	Nama      string `json:"nama" valid:"required~field nama tidak boleh kosong"`
	Kaprodi   string `json:"kaprodi" valid:"required~field kaprodi tidak boleh kosong"`
	JurusanID string `json:"jurusan" valid:"required~field jurusan tidak boleh kosong"`
}
