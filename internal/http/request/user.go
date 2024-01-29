package request

type SignIn struct {
	Username string `json:"username" binding:"required~field credential tidak ditemukan"`
	Password string `json:"password" binding:"required~field credential tidak ditemukan"`
}
