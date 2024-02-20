package response

type TeacherAttendence struct {
	Uuid         string        `json:"uuid"`
	Middle       string        `json:"middle"`
	Last         string        `json:"last"`
	GeadeOnTime  bool          `json:"grade_on_time"`
	Teacher      *Teacher      `json:"teacher"`
	Subject      *Subject      `json:"subject"`
	AcademicYear *AcademicYear `json:"academic_year"`
	CreatedAt    int64         `json:"created_at"`
	UpdatedAt    int64         `json:"updated_at"`
}

type StudentAttendence struct {
	Uuid          string        `json:"uuid"`
	StudentAmount string        `json:"student_amount"`
	Middle        string        `json:"middle"`
	Last          string        `json:"last"`
	PassedAmount  string        `json:"passed_amount"`
	FinalAmount   string        `json:"final_amount"`
	Subject       *Subject      `json:"subject"`
	AcademicYear  *AcademicYear `json:"academic_year"`
	CreatedAt     int64         `json:"created_at"`
	UpdatedAt     int64         `json:"updated_at"`
}
