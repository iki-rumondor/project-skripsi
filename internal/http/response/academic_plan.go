package response

type AcademicPlan struct {
	Uuid         string        `json:"uuid"`
	Middle       bool          `json:"middle"`
	Last         bool          `json:"last"`
	Available    *bool         `json:"available"`
	Note         string        `json:"note"`
	Subject      *Subject      `json:"subject"`
	AcademicYear *AcademicYear `json:"academic_year"`
	CreatedAt    int64         `json:"created_at"`
	UpdatedAt    int64         `json:"updated_at"`
}

type Rps struct {
	Uuid         string        `json:"uuid"`
	Status       bool          `json:"status"`
	Accept       bool          `json:"accept"`
	Note         *string       `json:"note"`
	FileName     *string       `json:"file_name"`
	Subject      *Subject      `json:"subject"`
	AcademicYear *AcademicYear `json:"academic_year"`
	CreatedAt    int64         `json:"created_at"`
	UpdatedAt    int64         `json:"updated_at"`
}
