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
	/* Test Put - Will need to fix up function properly */
	api.Put("/users/update/:id", handlers.UpdateUser)

	/* Test for User Password - Will need to add proper functionality */
	/*
		api.Post("/register", handlers.CreateUserPass)
		api.Post("/login/:id", handlers.GetUserPass)
		api.Put("/users/updatepassphrase", handlers.UpdateUserPass)
	*/

	// Forum routes
	api.Post("/forums", handlers.CreateForum)
	api.Get("/forums/:id", handlers.GetForum)
}
