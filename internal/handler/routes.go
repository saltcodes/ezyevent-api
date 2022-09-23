package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func InitRoutes(app *fiber.App) {

	app.Use(cors.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello ezyevents API v1.0")
	})
}
