package main

import (
	"ezyevent-api/internal/database"
	"ezyevent-api/internal/handler"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func main() {
	app := fiber.New()
	database.InitDatabase()
	handler.InitRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Fatalln(app.Listen("0.0.0.0:" + port))
}
