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
	}
}
