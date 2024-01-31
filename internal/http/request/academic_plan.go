package request

type AcademicPlan struct {
	Name             string `json:"name" valid:"required~field name tidak ditemukan"`
	Available        bool   `json:"available" valid:"required~field available tidak ditemukan"`
	Note             string `json:"note" valid:"required~field note tidak ditemukan"`
	SubjectUuid      string `json:"subject" valid:"required~field subject_uuid tidak ditemukan"`
	AcademicYearUuid string `json:"academic_year" valid:"required~field academic_year_uuid tidak ditemukan"`
}
