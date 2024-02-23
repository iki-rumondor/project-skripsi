package request

type SignIn struct {
	Username string `json:"username" valid:"required~field username tidak ditemukan"`
	Password string `json:"password" valid:"required~field password tidak ditemukan"`
}

type StepMonev struct {
	ID   uint   `json:"id"`
	Step string `json:"step" valid:"required~field tahapan tidak ditemukan"`
}
