package response

type TeacherSkill struct {
	Uuid      string   `json:"uuid"`
	Skill     string   `json:"skill"`
	Teacher   *Teacher `json:"teacher"`
	Subject   *Subject `json:"subject"`
	CreatedAt int64    `json:"created_at"`
	UpdatedAt int64    `json:"updated_at"`
}
