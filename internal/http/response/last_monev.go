package response

type StudentPassed struct {
	Uuid          string        `json:"uuid"`
	StudentAmount string        `json:"student_amount"`
	PassedAmount  string        `json:"passed_amount"`
	Subject       *Subject      `json:"subject"`
	AcademicYear  *AcademicYear `json:"academic_year"`
	CreatedAt     int64         `json:"created_at"`
	UpdatedAt     int64         `json:"updated_at"`
}
