package request

type AcademicYear struct {
	FirstDate      string `json:"first_date" valid:"required~field tanggal awal semester tidak ditemukan"`
	FirstDays      string `json:"first_days" valid:"required~field jumlah hari awal semester tidak ditemukan, range(1|9999)~field jumlah minimal 1"`
	MiddleDate     string `json:"middle_date" valid:"required~field tanggal tengah semester tidak ditemukan"`
	MiddleDays     string `json:"middle_days" valid:"required~field jumlah hari tengah semester tidak ditemukan, range(1|9999)~field jumlah minimal 1"`
	MiddleLastDate string `json:"middle_last_date" valid:"required~field tanggal sebelum UAS tidak ditemukan"`
	MiddleLastDays string `json:"middle_last_days" valid:"required~field jumlah hari sebelum UAS tidak ditemukan, range(1|9999)~field jumlah minimal 1"`
	LastDate       string `json:"last_date" valid:"required~field tanggal setelah UAS tidak ditemukan"`
	LastDays       string `json:"last_days" valid:"required~field jumlah hari setelah UAS tidak ditemukan, range(1|9999)~field jumlah minimal 1"`
	Year           string `json:"year" valid:"required~field tahun tidak ditemukan"`
	Semester       string `json:"semester" valid:"required~field semester tidak ditemukan"`
}

type UpdateTimeMonev struct {
	FirstDate        string `json:"first_date"`
	FirstAmount      string `json:"first_amount"`
	MiddleDate       string `json:"middle_date"`
	MiddleAmount     string `json:"middle_amount"`
	MiddleLastDate   string `json:"middle_last_date"`
	MiddleLastAmount string `json:"middle_last_amount"`
	LastDate         string `json:"last_date"`
	LastAmount       string `json:"last_amount"`
}

type UpdateLast struct {
	LastDate   string `json:"last_date" valid:"required~field tanggal setelah UAS tidak ditemukan"`
	LastAmount string `json:"last_amount" valid:"required~field jumlah hari setelah UAS tidak ditemukan, range(1|9999)~field jumlah minimal 1"`
}

type UpdateMiddleLast struct {
	MiddleLastDate   string `json:"middle_last_date" valid:"required~field tanggal sebelum UAS tidak ditemukan"`
	MiddleLastAmount string `json:"middle_last_amount" valid:"required~field jumlah hari sebelum UAS tidak ditemukan, range(1|9999)~field jumlah minimal 1"`
}

type UpdateMiddle struct {
	MiddleDate   string `json:"middle_date" valid:"required~field tanggal tengah semester tidak ditemukan"`
	MiddleAmount string `json:"middle_amount" valid:"required~field jumlah hari tengah semester tidak ditemukan, range(1|9999)~field jumlah minimal 1"`
}

type UpdateFirst struct {
	FirstDate   string `json:"first_date" valid:"required~field tanggal awal semester tidak ditemukan"`
	FirstAmount string `json:"first_amount" valid:"required~field jumlah hari awal semester tidak ditemukan, range(1|9999)~field jumlah minimal 1"`
}

type UpdateAcademicYear struct {
	Year     string `json:"year" valid:"required~field tahun tidak ditemukan"`
	Semester string `json:"semester" valid:"required~field semester tidak ditemukan"`
}

type AcademicYearOpen struct {
	Open bool `json:"open" `
}
