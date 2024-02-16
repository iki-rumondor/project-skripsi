package response

type FacilityCondition struct {
	Uuid         string        `json:"uuid"`
	Amount       string        `json:"amount"`
	Unit         string        `json:"unit"`
	Deactive     string        `json:"deactive"`
	Note         string        `json:"note"`
	CreatedAt    int64         `json:"created_at"`
	UpdatedAt    int64         `json:"updated_at"`
	Facility     *Facility     `json:"facility"`
	AcademicYear *AcademicYear `json:"academic_year"`
}