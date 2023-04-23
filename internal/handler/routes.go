package handler

import (
	"ezyevent-api/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var db = &database.DBConn

func InitRoutes(app *fiber.App) {

	app.Use(cors.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello ezyevents API v1.0")
	})

	app.Get("/event-types", ListEventTypes)
	app.Post("/event-types", CreateEventType)
	app.Get("/users", GetUsers)
	app.Get("/comments", ListComments)
	app.Post("/events", CreateEvent)
}
