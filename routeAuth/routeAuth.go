package routeAuth

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func RouteAuth(c *fiber.Ctx) error {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error loading .env file"})
	}

	// Get the token from the Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization header is missing"})
	}

	// Split the Bearer token
	tokenStr := strings.Split(authHeader, " ")
	if len(tokenStr) != 2 || tokenStr[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
	}

	// Parse the token
	token, err := jwt.Parse(tokenStr[1], func(token *jwt.Token) (interface{}, error) {
		// Make sure the signing method is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	// Token is valid, proceed to the next handler
	return c.Next()
}
