package request

type Register struct {
	Username string `json:"username" valid:"required~field username is required"`
	Email    string `json:"email" valid:"required~field email is required, email"`
	Password string `json:"password" valid:"required~field password is required, length(6|99)~password at least 6 character"`
	RoleID   uint   `json:"role_id" valid:"required~field role_id is required"`
}

type Login struct {
	Email    string `json:"email" valid:"required~please make sure to provide password in the request, email"`
	Password string `json:"password" valid:"required~please make sure to provide password in the request "`
}
