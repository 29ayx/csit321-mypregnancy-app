package handlers

import (
    "context"
    "net/http"
    "time"
    "gofiber-mongodb/server/database"
    "gofiber-mongodb/models"
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User Payload"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func CreateUser(c *fiber.Ctx) error {
    collection := database.GetCollection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var user models.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
    }

    result, err := collection.InsertOne(ctx, user)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
    }

    return c.Status(http.StatusOK).JSON(result)
}

// GetUser godoc
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [get]
func GetUser(c *fiber.Ctx) error {
    collection := database.GetCollection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    id, _ := primitive.ObjectIDFromHex(c.Params("id"))
    var user models.User
    err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
    if err != nil {
        return c.Status(http.StatusNotFound).JSON(map[string]string{"error": "User not found"})
    }

    return c.Status(http.StatusOK).JSON(user)
}

// Add other handlers for update and delete with similar annotations
