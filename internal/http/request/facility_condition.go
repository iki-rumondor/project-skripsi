package request

type FacilityCondition struct {
	Note             string `json:"note" valid:"required~field keterangan tidak ditemukan"`
	Amount           string `json:"amount" valid:"required~field jumlah tidak ditemukan, range(1|9999)~field jumlah minimal 1"`
	Unit             string `json:"unit" valid:"required~field satuan tidak ditemukan"`
	Deactive         string `json:"deactive" valid:"required~field tidak berfungsi tidak ditemukan, range(0|9999)~field tidak berfungsi minimal 0"`
	FacilityUuid     string `json:"facility_uuid" valid:"required~field fasilitas tidak ditemukan"`
	AcademicYearUuid string `json:"academic_year_uuid" valid:"required~field tahun ajaran tidak ditemukan"`
}

type UpdateFacilityCondition struct {
	Note     string `json:"note" valid:"required~field keterangan tidak ditemukan"`
	Amount   string `json:"amount" valid:"required~field jumlah tidak ditemukan"`
	Unit     string `json:"unit" valid:"required~field jumlah tidak ditemukan"`
	Deactive string `json:"deactive" valid:"required~field tidak berfungsi tidak ditemukan, range(0|9999)~field tidak berfungsi minimal 0"`
}
