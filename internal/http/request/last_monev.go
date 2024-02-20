package request

type UpdateStudentPass struct {
	StudentAmount string `json:"student_amount" valid:"required~field jumlah mahasiswa tidak ditemukan, range(1|9999)~field jumlah mahasiswa minimal 1"`
	PassedAmount  string `json:"passed_amount" valid:"required~field jumlah mahasiswa lulus tidak ditemukan, range(1|9999)~field jumlah mahasiswa lulus minimal 1"`
}

type UpdateStudentFinal struct {
	StudentAmount string `json:"student_amount" valid:"required~field jumlah mahasiswa tidak ditemukan, range(1|9999)~field jumlah mahasiswa minimal 1"`
	FinalAmount   string `json:"final_amount" valid:"required~field jumlah mahasiswa mengikuti UAS tidak ditemukan, range(1|9999)~field jumlah mahasiswa mengikuti UAS minimal 1"`
}

type UpdateTeacherGrade struct {
	GradeOnTime bool `json:"grade"`
}
