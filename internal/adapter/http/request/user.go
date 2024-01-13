package request

type Register struct {
	Username string `json:"username" valid:"required~field username is required"`
	Email    string `json:"email" valid:"required~field email is required, email"`
	Password string `json:"password" valid:"required~field password is required, length(6|99)~password at least 6 character"`
	Role     string `json:"rolrolee_id" valid:"required~field role is required"`
}

type Login struct {
	Username string `json:"username" valid:"required~field username is required"`
	Password string `json:"password" valid:"required~field password is required "`
}
