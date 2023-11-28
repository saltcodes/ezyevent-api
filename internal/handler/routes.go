package handler

import (
	"ezyevent-api/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var db = &database.DBConn

// Initiate gRPC client Connection
var con, grpcErr = grpc.Dial("localhost:8181", grpc.WithTransportCredentials(insecure.NewCredentials()))

func InitRoutes(app *fiber.App) {

	app.Use(cors.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello ezyevents API v1.0")
	})

	//Events End point
	app.Get("/events", ListEvents)
	app.Post("/events", CreateEvent)
	app.Get("/events/find", FindEvents)
	app.Get("/events/:id", GetEvent)
	app.Put("/events/:id", UpdateEvent)
	app.Delete("/events/:id", DeleteEvent)

	//Event type endpoints
	app.Get("/event-types", ListEventTypes)
	app.Post("/event-types", CreateEventType)

	//Users
	app.Get("/users", GetUsers)
	app.Post("/users", CreateUser)
	app.Get("/users/:id", GetUser)
	app.Delete("users/:id", DeleteUser)
	app.Put("/users/:id", UpdateUser)

	//Comments endpoint
	app.Get("/comments", ListComments)

	//EventSchedule
	app.Get("/event-schedules", ListEventSchedule)
	app.Get("/event-schedules/:id", GetEventSchedule)
	app.Put("/event-schedules/:id", UpdateEventSchedule)
	app.Post("/event-schedules/", CreateEventSchedule)
	app.Delete("/event-schedules", DeleteEventSchedule)
}
