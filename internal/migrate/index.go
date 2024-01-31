package migrate

import (
	"fmt"
	"log"

	"github.com/iki-rumondor/go-monev/internal/models"
	"gorm.io/gorm"
)

func ReadTerminal(db *gorm.DB, args []string) {
	switch {
	case args[1] == "fresh":
		if err := freshDatabase(db); err != nil {
			log.Fatal(err.Error())
		}
	case args[1] == "migrate":
		if err := migrateDatabase(db); err != nil {
			log.Fatal(err.Error())
		}
	default:
		fmt.Println("Hello, Nice to meet you")
	}
}

func freshDatabase(db *gorm.DB) error {
	for _, model := range GetAllModels() {
		if err := db.Migrator().DropTable(model.Model); err != nil {
			return err
		}
	}
	for _, model := range GetAllModels() {
		if err := db.Debug().AutoMigrate(model.Model); err != nil {
			return err
		}
	}

	db.Create(&models.Role{
		Name: "ADMIN",
		User: &models.User{
			Username: "admin",
			Password: "123",
		},
	})

	db.Create(&models.Role{
		Name: "DEPARTMENT",
	})

	return nil
}

func migrateDatabase(db *gorm.DB) error {
	for _, model := range GetAllModels() {
		if err := db.Debug().AutoMigrate(model.Model); err != nil {
			return err
		}
	}
	return nil
}
