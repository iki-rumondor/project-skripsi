package main

import (
	"log"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/database"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/registry"
	"github.com/iki-rumondor/init-golang-service/internal/routes"
	"gorm.io/gorm"
)

func main() {
	gormDB, err := database.NewMysqlDB()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// freshDatabase(gormDB)

	for _, model := range registry.RegisterModels() {
		if err := gormDB.Debug().AutoMigrate(model.Model); err != nil {
			log.Fatal(err.Error())
			return
		}
	}

	repo := registry.GetRepositories(gormDB)
	handlers := registry.GetHandlers(repo)

	// dbSeeder(gormDB)

	var PORT = ":8080"
	routes.StartServer(handlers).Run(PORT)
}

func dbSeeder(db *gorm.DB) {

	db.Create(&domain.User{
		Username: "admin",
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

func freshDatabase(db *gorm.DB) {
	for _, model := range registry.RegisterModels() {
		if err := db.Migrator().DropTable(model.Model); err != nil {
			log.Fatal(err.Error())
			return
		}
	}
}
