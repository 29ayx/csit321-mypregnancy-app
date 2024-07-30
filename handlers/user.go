package handlers

import (
	"context"
	"gofiber-mongodb/models"
	"gofiber-mongodb/server/database"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
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

	// Validate password length
	if len(user.PassHash) < 8 {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": "Password must be at least 8 characters long"})
	}

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PassHash), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": "Failed to hash password"})
	}

	// Store the hashed password instead of the plain password
	user.PassHash = string(hashedPassword)

	// Check if the email is already taken
	count, err := collection.CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": "Database error"})
	}

	if count > 0 {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": "Email is already taken"})
	}

	// Insert the user into the database
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}

	// Generate JWT token for the new user
	token, err := GenerateToken(user.Email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": "Failed to generate token"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"userid": result.InsertedID,
		"token":  token,
	})
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

func LoginUser(c *fiber.Ctx) error {
	var loginRequest struct {
		Email    string `json:"email"`
		PassHash string `json:"passhash"` // Changed from "Password" to "PassHash"
	}

	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}

	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err := collection.FindOne(ctx, bson.M{"email": loginRequest.Email}).Decode(&user)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(map[string]string{"error": "Invalid email or password"})
	}

	// Verify the password using bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(loginRequest.PassHash)); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(map[string]string{"error": "Invalid email or password"})
	}

	return c.Status(http.StatusOK).JSON(map[string]string{"message": "Login successful"})
}

// GenerateToken generates a JWT token for the given user ID and email
func GenerateToken(email string) (string, error) {
	// Create a new token object, specifying signing method and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 48).Unix(), // Token expiration time
	})

	// Sign and get the complete encoded token as a string
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
