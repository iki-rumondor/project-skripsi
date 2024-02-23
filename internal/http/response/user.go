package response

type User struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type SubjectsCount struct {
	General   int `json:"general"`
	Practical int `json:"practical"`
}

type MonevAmount struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type Setting struct {
	ID    interface{} `json:"id"`
	Name  interface{} `json:"name"`
	Value interface{} `json:"value"`
}
