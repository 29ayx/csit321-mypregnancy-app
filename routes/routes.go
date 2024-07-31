package routes

import (
	"gofiber-mongodb/handlers"
	//"gofiber-mongodb/routeAuth"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// User routes
	api.Post("/signup", handlers.CreateUser)
	api.Post("/login", handlers.LoginUser)
	api.Get("/users/:id", handlers.GetUser)
	api.Put("/users/update/:id", handlers.UpdateUser)

	// Forum routes
	api.Post("/forums", handlers.CreateForum)
	api.Get("/forums/:id", handlers.GetForum)
}
