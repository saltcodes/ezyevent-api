package main

import (
	"ezyevent-api/internal/handler"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func main() {
	app := fiber.New()
	handler.InitRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Fatalln(app.Listen(":" + port))
}
