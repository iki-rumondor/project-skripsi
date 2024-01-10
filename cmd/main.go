package main

import (
	"log"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/database"
	customHTTP "github.com/iki-rumondor/init-golang-service/internal/adapter/http"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/registry"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"github.com/iki-rumondor/init-golang-service/internal/routes"
	"gorm.io/gorm"
)

func main() {
	gormDB, err := database.NewMysqlDB()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	for _, model := range registry.RegisterModels() {
		if err := gormDB.Debug().AutoMigrate(model.Model); err != nil {
			log.Fatal(err.Error())
			return
		}
	}

	// dbSeeder(gormDB)

	auth_repo := repository.NewAuthRepository(gormDB)
	auth_service := application.NewAuthService(auth_repo)
	auth_handler := customHTTP.NewAuthHandler(auth_service)

	util_repo := repository.NewUtilRepository(gormDB)
	util_service := application.NewUtilService(util_repo)
	util_handler := customHTTP.NewUtilHandler(util_service)

	prodi_repo := repository.NewProdiRepository(gormDB)
	prodi_service := application.NewProdiService(prodi_repo)
	prodi_handler := customHTTP.NewProdiHandler(prodi_service)

	instrumen_repo := repository.NewInstrumenRepository(gormDB)
	admin_service := application.NewAdminService(repository.Repositories{
		InstrumenRepository: instrumen_repo,
	})
	admin_handler := customHTTP.NewAdminHandler(admin_service)

	handlers := &customHTTP.Handlers{
		AuthHandler:  auth_handler,
		UtilHandler:  util_handler,
		ProdiHandler: prodi_handler,
		AdminHandler: admin_handler,
	}

	var PORT = ":8080"
	routes.StartServer(handlers).Run(PORT)
}

func dbSeeder(db *gorm.DB) {

	db.Create(&domain.User{
		Uuid:     "useruuid",
		Username: "admin",
		Email:    "admin@admin.com",
		Password: "123",
		Role:     "ADMIN",
	})

	db.Create(&domain.Jurusan{
		Nama: "Teknik Informatika",
	})
	db.Create(&domain.Jurusan{
		Nama: "Teknik Sipil",
	})
	db.Create(&domain.Jurusan{
		Nama: "Teknik Arsitektur",
	})
}
