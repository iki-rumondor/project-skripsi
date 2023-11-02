package main

import (
	"log"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/database"
	customHTTP "github.com/iki-rumondor/init-golang-service/internal/adapter/http"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"github.com/iki-rumondor/init-golang-service/internal/routes"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
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

	handlers := &customHTTP.Handlers{
		AuthHandler: auth_handler,
	}

	utils.NewCustomValidator(gormDB)
	var PORT = ":8080"
	routes.StartServer(handlers).Run(PORT)
}

func autoMigration(db *gorm.DB) {

	db.Migrator().DropTable(&domain.Role{})
	db.Migrator().DropTable(&domain.User{})
	db.Migrator().CreateTable(&domain.Role{})
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
}
