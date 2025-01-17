package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"gofiber-mongodb/models"
	"gofiber-mongodb/server/database"
	"log"
	"net/http"
	"os"
	"strings"
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
	// Extract the token from the Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(http.StatusUnauthorized).JSON(map[string]string{"error": "Missing or malformed JWT"})
	}

	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return c.Status(http.StatusUnauthorized).JSON(map[string]string{"error": "Malformed JWT"})
	}
	tokenString := splitToken[1]

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(map[string]string{"error": "Invalid or expired JWT"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.Status(http.StatusUnauthorized).JSON(map[string]string{"error": "Invalid or expired JWT"})
	}

	// Check if the "email" claim is present and is a string
	email, ok := claims["email"].(string)
	if !ok || email == "" {
		return c.Status(http.StatusUnauthorized).JSON(map[string]string{"error": "Invalid JWT claims"})
	}

	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err = collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(map[string]string{"error": "User not found"})
	}

	// Strip passhash from response
	user.PassHash = ""

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

	var requestData struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}
	
	if err := c.BodyParser(&requestData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}

	// Validate password length
	if len(requestData.Password) < 8 {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": "Invalid password length"})
	}

	// Check if the email is already taken
	count, err := collection.CountDocuments(ctx, bson.M{"email": requestData.Email})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": "Database error"})
	}

	if count > 0 {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": "Email is already taken"})
	}

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestData.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": "Failed to hash password"})
	}

	// Create the user object with the hashed password
	user := models.User{
		FirstName: requestData.FirstName,
		LastName:  requestData.LastName,
		Email:     requestData.Email,
		PassHash:  string(hashedPassword),
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
		Password string `json:"password"`
	}

	// Manually parse the JSON body
	if err := json.Unmarshal(c.Body(), &loginRequest); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}

	// Validate password length
	if len(loginRequest.Password) < 8 {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": "Invalid email or password"})
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
	err = bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(loginRequest.Password))
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(map[string]string{"error": "Invalid email or password"})
	}

	// Generate JWT token
	token, err := GenerateToken(user.Email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": "Failed to generate token"})
	}

	return c.Status(http.StatusOK).JSON(map[string]string{
		"message": "Login successful",
		"token":   token,
	})
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
