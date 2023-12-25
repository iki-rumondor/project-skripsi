package response

type Message struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}
