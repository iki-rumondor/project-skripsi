package request

type Facility struct {
	Name string `json:"name" valid:"required~field name tidak ditemukan"`
}
