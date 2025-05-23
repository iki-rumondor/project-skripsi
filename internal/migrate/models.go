package migrate

import "github.com/iki-rumondor/go-monev/internal/models"

type Model struct {
	Model interface{}
}

func GetAllModels() []Model {
	return []Model{
		{Model: models.Role{}},
		{Model: models.User{}},
		{Model: models.Department{}},
		{Model: models.Subject{}},
		{Model: models.Major{}},
		{Model: models.Laboratory{}},
		{Model: models.AcademicYear{}},
		{Model: models.AcademicPlan{}},
		{Model: models.PracticalModule{}},
		{Model: models.PracticalTool{}},
		{Model: models.Teacher{}},
		{Model: models.TeacherSkill{}},
		{Model: models.Facility{}},
		{Model: models.FacilityCondition{}},
		{Model: models.TeacherAttendence{}},
		{Model: models.StudentAttendence{}},
		{Model: models.Setting{}},
		{Model: models.Rps{}},
	}
}
