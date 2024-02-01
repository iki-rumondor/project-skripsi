package response

type Subject struct {
	Uuid         string        `json:"uuid"`
	Name         string        `json:"name"`
	Code         string        `json:"code"`
	Practical    *bool          `json:"practical"`
	CreatedAt    int64         `json:"created_at"`
	UpdatedAt    int64         `json:"updated_at"`
	Department   *Department   `json:"department"`
	AcademicPlan *AcademicPlan `json:"academic_plan"`
}
