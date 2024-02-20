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
