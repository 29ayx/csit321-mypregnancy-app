package handlers

import (
	"context"
	"gofiber-mongodb/models"
	"gofiber-mongodb/server/database"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

// Global UpdateUserField to be referenced
func UpdateUser(c *fiber.Ctx) error {
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Parse the user ID from the request parameters
	userID := c.Params("id")
	if userID == "" {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": "User ID is required"})
	}

	// Parse the request body for the fields to be updated
	var updateData map[string]interface{}
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}

	// Ensure that the updateData is not empty
	if len(updateData) == 0 {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": "No update data provided"})
	}

	// Convert the user ID to an ObjectID
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": "Invalid user ID format"})
	}

	// Create the update document
	update := bson.M{
		"$set": updateData,
	}

	// Perform the update operation
	result, err := collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}

	// Check if any document was modified
	if result.ModifiedCount == 0 {
		return c.Status(http.StatusNotFound).JSON(map[string]string{"error": "User not found or no changes made"})
	}

	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"message":       "User updated successfully",
		"modifiedCount": result.ModifiedCount,
	})
}
