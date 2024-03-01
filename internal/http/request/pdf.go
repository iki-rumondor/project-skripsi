package request

type PdfReport struct {
	Department string      `json:"department"`
	Laboratory string      `json:"laboratory"`
	Semester   string      `json:"semester"`
	Year       string      `json:"year"`
	Data       interface{} `json:"data"`
}

type PdfPlans struct {
	Subject string `json:"subject"`
	Status  string `json:"status"`
	Note    string `json:"note"`
}
