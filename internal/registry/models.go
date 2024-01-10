package registry

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: domain.Indikator{}},
		{Model: domain.IndikatorType{}},
		{Model: domain.InstrumenType{}},
		{Model: domain.InstrumenType{}},
		{Model: domain.Jurusan{}},
		{Model: domain.Prodi{}},
		{Model: domain.Role{}},
		{Model: domain.User{}},
	}
}
