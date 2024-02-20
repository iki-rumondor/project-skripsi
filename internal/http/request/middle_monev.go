package request

type CreateTeacherAttendence struct {
	SubjectUuid      string `json:"subject_uuid" valid:"required~field mata kuliah tidak ditemukan"`
	AcademicYearUuid string `json:"academic_year_uuid" valid:"required~field tahun ajaran tidak ditemukan"`
	Middle           string `json:"middle" valid:"required~field middle tidak ditemukan"`
}

type CreateStudentAttendence struct {
	SubjectUuid      string `json:"subject_uuid" valid:"required~field mata kuliah tidak ditemukan"`
	AcademicYearUuid string `json:"academic_year_uuid" valid:"required~field tahun ajaran tidak ditemukan"`
	StudentAmount    string `json:"student_amount" valid:"required~field jumlah mahasiswa tidak ditemukan, int~field jumlah mahasiswa tidak valid"`
	Middle           string `json:"middle" valid:"required~field jumlah kehadiran mahasiswa tidak ditemukan, int~field jumlah kehadiran mahasiswa tidak valid"`
}

type LastTeacherAttendence struct {
	Last string `json:"last" valid:"required~field jumlah kehadiran mahasiswa tidak ditemukan, int~field jumlah kehadiran mahasiswa tidak valid"`
}

type LastStudentAttendence struct {
	Last string `json:"last" valid:"required~field persentase tidak ditemukan, int~field persentase tidak valid"`
}
