package response

type DataUsers struct {
	Success bool    `json:"success"`
	Data    []*User `json:"data"`
}

type User struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
