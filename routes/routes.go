package routes

import (
	"gofiber-mongodb/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// User routes
	api.Post("/users", handlers.CreateUser)
	api.Get("/users/:id", handlers.GetUser)
	/* Test Put - Will need to test updating each field in front end */
	api.Put("/users/update/:id", handlers.UpdateFirstName)

	// Forum routes
	api.Post("/forums", handlers.CreateForum)
	api.Get("/forums/:id", handlers.GetForum)
}
