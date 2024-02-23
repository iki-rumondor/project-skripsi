package request

type AcademicYear struct {
	Year     string `json:"year" valid:"required~field tahun tidak ditemukan"`
	Semester string `json:"semester" valid:"required~field semester tidak ditemukan"`
}

type AcademicYearOpen struct {
	Open bool `json:"open" `
}
