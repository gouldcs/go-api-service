package main

import (
	"gouldcs/gouldservice/internal/app"
	"gouldcs/gouldservice/internal/config"
	"log"
)

func main() {
	db, err := config.GetDB()

	if err != nil {
		log.Fatal("Failed to start the database.")
	}

	router := app.SetupRouter(db)

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}