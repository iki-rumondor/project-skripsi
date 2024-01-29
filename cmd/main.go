package main

import (
	"fmt"
	"log"
	"os"

	"github.com/iki-rumondor/go-monev/internal/config"
	"github.com/iki-rumondor/go-monev/internal/http/routes"
)

func main() {
	gormDB, err := config.NewMysqlDB()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	handlers := config.GetAppHandlers(gormDB)

	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	routes.StartServer(handlers).Run(fmt.Sprintf(":%s", PORT))
}
