package handler

import (
	"ezyevent-api/internal/database"
	"ezyevent-api/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var db = &database.DBConn

// Initiate gRPC client Connection
var con, grpcErr = grpc.Dial("localhost:8181", grpc.WithTransportCredentials(insecure.NewCredentials()))

func InitRoutes(app *fiber.App) {

	app.Use(cors.New())

	v1 := app.Group("/v1")

	v1.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello ezyevents API v1.0")
	})

	//vysCiv-najte3-gyjjoh

	unRestrictedRoutes := v1.Group("/")
	//Restricted Routes with Auth0  for CUD

	restrictedRoutes := v1.Group("/")
	//Middleware
	restrictedRoutes.Use(adaptor.HTTPMiddleware(middleware.EnsureValidToken()))

	//Events Endpoint
	unRestrictedRoutes.Get("/events", ListEvents)
	unRestrictedRoutes.Get("/events/find", FindEvents)
	unRestrictedRoutes.Get("/events/:id", GetEvent)

	restrictedRoutes.Put("/events/:id", UpdateEvent)
	restrictedRoutes.Delete("/events/:id", DeleteEvent)
	restrictedRoutes.Post("/events", CreateEvent)

	//Categories Endpoint
	restrictedRoutes.Post("/categories", CreateCategory)
	unRestrictedRoutes.Get("/categories", ListCategories)

	//Users Endpoint
	unRestrictedRoutes.Get("/users", GetUsers)
	unRestrictedRoutes.Get("/users/:id", GetUser)
	restrictedRoutes.Delete("users/:id", DeleteUser)
	restrictedRoutes.Put("/users/:id", UpdateUser)
	restrictedRoutes.Post("/users", CreateUser)

	//Schedules Endpoint
	unRestrictedRoutes.Get("/schedules", ListSchedule)
	unRestrictedRoutes.Get("/schedules/:id", GetSchedule)
	restrictedRoutes.Put("/schedules/:id", UpdateSchedule)
	restrictedRoutes.Post("/schedules/", CreateSchedule)
	restrictedRoutes.Delete("/schedules", DeleteEventSchedule)

	//Organizations Endpoint
	unRestrictedRoutes.Get("/organizations", ListOrganization)
	unRestrictedRoutes.Get("/organizations/:id", GetOrganization)
	restrictedRoutes.Put("/organizations/:id", UpdateOrganization)
	restrictedRoutes.Post("/organizations", CreateOrganization)
	restrictedRoutes.Delete("/organizations/:id", DeleteOrganization)

	//Booking endpoint
	unRestrictedRoutes.Get("/bookings", ListBooking)
	unRestrictedRoutes.Get("/bookings/:id", GetBooking)
	restrictedRoutes.Put("/bookings/:id", UpdateBooking)
	restrictedRoutes.Post("/bookings", CreateBooking)
	restrictedRoutes.Delete("/bookings/:id", DeleteBooking)

	//Comments endpoint
	unRestrictedRoutes.Get("/comments", ListComments)
}
