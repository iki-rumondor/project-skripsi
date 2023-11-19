package main

import (
	"log"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/database"
	customHTTP "github.com/iki-rumondor/init-golang-service/internal/adapter/http"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
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
	
	// autoMigration(gormDB)

	auth_repo := repository.NewAuthRepository(gormDB)
	auth_service := application.NewAuthService(auth_repo)
	auth_handler := customHTTP.NewAuthHandler(auth_service)

	util_repo := repository.NewUtilRepository(gormDB)
	util_service := application.NewUtilService(util_repo)
	util_handler := customHTTP.NewUtilHandler(util_service)

	prodi_repo := repository.NewProdiRepository(gormDB)
	prodi_service := application.NewProdiService(prodi_repo)
	prodi_handler := customHTTP.NewProdiHandler(prodi_service)

	handlers := &customHTTP.Handlers{
		AuthHandler: auth_handler,
		UtilHandler: util_handler,
		ProdiHandler: prodi_handler,
	}

	var PORT = ":8080"
	routes.StartServer(handlers).Run(PORT)
}

func autoMigration(db *gorm.DB) {

	db.Migrator().DropTable(&domain.Role{})
	db.Migrator().CreateTable(&domain.Role{})

	db.Migrator().DropTable(&domain.Jurusan{})
	db.Migrator().CreateTable(&domain.Jurusan{})

	db.Migrator().DropTable(&domain.Prodi{})
	db.Migrator().CreateTable(&domain.Prodi{})

	db.Migrator().DropTable(&domain.User{})
	db.Migrator().CreateTable(&domain.User{})

	roles := []domain.Role{
		{
			Name: "Admin",
		},
		{
			Name: "User",
		},
	}

	db.Create(&roles)

	db.Create(&domain.User{
		Uuid: "useruuid",
		Username: "admin",
		Email: "admin@admin.com",
		Password: "123",
		RoleID: 1,
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
