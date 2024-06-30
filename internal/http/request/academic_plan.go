package request

type AcademicPlan struct {
	Available        bool   `form:"available"`
	Note             string `form:"note"`
	SubjectUuid      string `form:"subject_uuid" valid:"required~field subject_uuid tidak ditemukan"`
	AcademicYearUuid string `form:"academic_year_uuid" valid:"required~field academic_year_uuid tidak ditemukan"`
}

type UpdateAcademicPlan struct {
	Status bool   `form:"status"`
	Note   string `form:"note"`
}

type AcademicPlanMiddle struct {
	Middle bool `json:"middle"`
}

type AcademicPlanLast struct {
	Last bool `json:"last"`
}
