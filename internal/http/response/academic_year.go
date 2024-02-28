package response

type AcademicYear struct {
	Uuid            string             `json:"uuid"`
	Name            string             `json:"name"`
	FirstDate       string             `json:"first_date"`
	MiddleDate      string             `json:"middle_date"`
	MiddleLastDate  string             `json:"middle_last_date"`
	LastDate        string             `json:"last_date"`
	FirstDays       string             `json:"first_days"`
	MiddleDays      string             `json:"middle_days"`
	MiddleLastDays  string             `json:"middle_last_days"`
	LastDays        string             `json:"last_days"`
	FirstRange      string             `json:"first_range"`
	MiddleRange     string             `json:"middle_range"`
	MiddleLastRange string             `json:"middle_last_range"`
	LastRange       string             `json:"last_range"`
	Year            string             `json:"year"`
	Semester        string             `json:"semester"`
	Status          string             `json:"status"`
	Open            bool               `json:"open"`
	AcademicPlan    *[]AcademicPlan    `json:"academic_plans"`
	PracticalTool   *[]PracticalTool   `json:"practical_tools"`
	PracticalModule *[]PracticalModule `json:"practical_modules"`
	CreatedAt       int64              `json:"created_at"`
	UpdatedAt       int64              `json:"updated_at"`
}
