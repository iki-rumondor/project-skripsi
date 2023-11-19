package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type UtilRepository interface{
	FindAllJurusan() (*[]domain.Jurusan, error)
}