package handler

import (
	"ezyevent-api/internal/database"
	"ezyevent-api/internal/middleware"
	"ezyevent-api/internal/util"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var db = &database.DBConn

// Initiate gRPC client Connection
var con, grpcErr = grpc.Dial(util.GetVariableWith("GRPC_HOST")+":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))

//	@title			EzyEvent Main API
//	@version		1.0
//	@description	Main API for CRUD Operations.
//	@contact.name	API Support
//	@contact.email	oseiyeboahjohnson@gmail.com
//
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath	/v1
func InitPrivateRoutes(app *fiber.App) {

	route := app.Group("/v1")

	route.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello ezyevents API v1.0")
	})

	//Restricted Routes with Auth0  for CUD
	//Middleware
	route.Use(adaptor.HTTPMiddleware(middleware.EnsureValidToken()))

	route.Put("/events/:id", UpdateEvent)
	route.Delete("/events/:id", DeleteEvent)
	route.Post("/events", CreateEvent)

	//Categories Endpoint
	route.Post("/categories", CreateCategory)

	//Users Endpoint
	route.Delete("users/:id", DeleteUser)
	route.Put("/users/:id", UpdateUser)
	route.Post("/users", CreateUser)

	//Schedules Endpoint
	route.Put("/schedules/:id", UpdateSchedule)
	route.Post("/schedules/", CreateSchedule)
	route.Delete("/schedules", DeleteEventSchedule)

	//Organizations Endpoint
	route.Put("/organizations/:id", UpdateOrganization)
	route.Post("/organizations", CreateOrganization)
	route.Delete("/organizations/:id", DeleteOrganization)

	//Booking endpoint
	route.Put("/bookings/:id", UpdateBooking)
	route.Post("/bookings", CreateBooking)
	route.Delete("/bookings/:id", DeleteBooking)

}

func InitPublicRoutes(app *fiber.App) {

	route := app.Group("/v1")
	//Events Endpoint
	route.Get("/events", ListEvents)
	route.Get("/events/find", FindEvents)
	route.Get("/events/:id", GetEvent)

	route.Get("/categories", ListCategories)

	//Users Endpoint
	route.Get("/users", GetUsers)
	route.Get("/users/:id", GetUser)

	//Schedules Endpoint
	route.Get("/schedules", ListSchedule)
	route.Get("/schedules/:id", GetSchedule)

	//Organizations Endpoint
	route.Get("/organizations", ListOrganization)
	route.Get("/organizations/:id", GetOrganization)

	//Booking endpoint
	route.Get("/bookings", ListBooking)
	route.Get("/bookings/:id", GetBooking)

	//Comments endpoint
	route.Get("/comments", ListComments)
}

func InitDocsRoute(app *fiber.App) {

	app.Use(swagger.New(swagger.Config{
		BasePath: "/v1",
		Title:    "EzyEvent API",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
	}))
}

// NotFoundRoute func for describe 404 Error route.
func NotFoundRoute(a *fiber.App) {
	// Register new special route.
	a.Use(
		// Anonimus function.
		func(c *fiber.Ctx) error {
			// Return HTTP 404 status and JSON response.
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"code":    true,
				"message": "sorry, endpoint is not found",
				"data":    nil,
			})
		},
	)
}
