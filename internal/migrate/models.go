package migrate

import "github.com/iki-rumondor/go-monev/internal/models"

type Model struct {
	Model interface{}
}

func GetAllModels() []Model {
	return []Model{
		{Model: models.User{}},
		{Model: models.Subject{}},
	}
}
