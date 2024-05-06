package main

import (
	"log"
	"github.com/epidemIT/epidemIT-Backend/database"
	"github.com/epidemIT/epidemIT-Backend/database/migrations"
	"github.com/epidemIT/epidemIT-Backend/route"

	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}

	database.InitDatabase()
	migrations.RunMigrations()

	app := fiber.New()

	route.SetupRoutes(app)

	app.Listen(":8080")
}
