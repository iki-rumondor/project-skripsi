package request

type CreateTeacherAttendence struct {
	SubjectUuid      string `json:"subject_uuid" valid:"required~field mata kuliah tidak ditemukan"`
	AcademicYearUuid string `json:"academic_year_uuid" valid:"required~field tahun ajaran tidak ditemukan"`
	TeacherUuid      string `json:"teacher_uuid" valid:"required~field penanggung jawab mata kuliah tidak ditemukan"`
	Middle           string `json:"middle" valid:"required~field middle tidak ditemukan"`
	Class            string `json:"class" valid:"required~field kelas kehadiran mahasiswa tidak ditemukan, stringlength(1|2)~field kelas tidak valid"`
}

type CreateStudentAttendence struct {
	SubjectUuid      string `json:"subject_uuid" valid:"required~field mata kuliah tidak ditemukan"`
	AcademicYearUuid string `json:"academic_year_uuid" valid:"required~field tahun ajaran tidak ditemukan"`
	StudentAmount    string `json:"student_amount" valid:"required~field jumlah mahasiswa tidak ditemukan, range(1|9999)~field jumlah minimal 1"`
	Middle           string `json:"middle" valid:"required~field jumlah kehadiran mahasiswa tidak ditemukan, range(1|9999)~field jumlah minimal 1"`
	Class            string `json:"class" valid:"required~field kelas kehadiran mahasiswa tidak ditemukan, stringlength(1|2)~field kelas tidak valid"`
}

type LastTeacherAttendence struct {
	Last string `json:"last" valid:"required~field jumlah kehadiran mahasiswa tidak ditemukan, range(1|9999)~field jumlah minimal 1"`
}

type LastStudentAttendence struct {
	Last string `json:"last" valid:"required~field persentase tidak ditemukan, range(1|9999)~field jumlah minimal 1"`
}
