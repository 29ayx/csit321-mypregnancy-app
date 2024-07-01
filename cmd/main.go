package main

import (
    "log"
    "gofiber-mongodb/server/database"
    "gofiber-mongodb/routes"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    _ "gofiber-mongodb/docs" // swagger docs
    "github.com/gofiber/swagger" // swagger middleware
)

// @title GoFiber MongoDB API
// @version 1.0
// @description This is a sample GoFiber server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /api
func main() {
    app := fiber.New()
    app.Use(logger.New())

    database.ConnectDB()
    routes.SetupRoutes(app)

    // Swagger route
    app.Get("/swagger/*", swagger.HandlerDefault)

    log.Fatal(app.Listen(":3000"))
}
