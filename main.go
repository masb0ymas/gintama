package main

import (
	"gintama/src/models"
	"gintama/src/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	// Initial Database
	models.ConnectDatabase()

	// Initial Routes
	r := routes.SetupRoutes()

	// Running App
	r.Run(":" + port)
}
