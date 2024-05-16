package main

import (
	"github.com/epidemIT/epidemIT-Backend/database"
	"github.com/epidemIT/epidemIT-Backend/database/migrations"
	"github.com/epidemIT/epidemIT-Backend/route"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.InitDatabase()
	migrations.RunMigrations()

	app := fiber.New()

	route.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}
