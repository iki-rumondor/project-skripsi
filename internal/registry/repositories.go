package registry

import (
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"gorm.io/gorm"
)

func GetRepositories(db *gorm.DB) *repository.Repositories {
	return &repository.Repositories{
		AuthRepo:      repository.NewAuthRepository(db),
		InstrumenRepo: repository.NewInstrumenRepository(db),
		ProdiRepo:     repository.NewProdiRepository(db),
		UtilRepo:      repository.NewUtilRepository(db),
	}
}
