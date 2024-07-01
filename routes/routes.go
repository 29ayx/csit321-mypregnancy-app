package routes

import (
	"gofiber-mongodb/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/users", handlers.CreateUser)
	api.Get("/users/:id", handlers.GetUser)

	// Forum routes
	api.Post("/forums", handlers.CreateForum)
	api.Get("/forums/:id", handlers.GetForum)

}
