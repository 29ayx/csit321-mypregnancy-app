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

func GetUserPass(c *fiber.Ctx) error {
	collection := database.GetCollection("userpassword")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, _ := primitive.ObjectIDFromHex(c.Params("id"))
	var pass models.UserPassword
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&pass)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(map[string]string{"error": "Password Incorrect"})
	}

	return c.Status(http.StatusOK).JSON(pass)
}

func CreateUserPass() {

	return
}

func UpdateUserPass() {

	return
}
