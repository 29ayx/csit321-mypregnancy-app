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

// CreateProfessionalAddress godoc
// @Summary Create a new professional address
// @Description Create a new professional address for a professional
// @Tags professionalAddresses
// @Accept  json
// @Produce  json
// @Param address body models.ProfessionalAddress true "Professional Address Payload"
// @Success 200 {object} models.ProfessionalAddress
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /professionalAddresses [post]
func CreateProfessionalAddress(c *fiber.Ctx) error {
	collection := database.GetCollection("professionalAddresses")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var address models.ProfessionalAddress
	if err := c.BodyParser(&address); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}

	address.ProfID = int(primitive.NewObjectID().Timestamp().Unix())

	_, err := collection.InsertOne(ctx, address)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(address)
}

// GetProfessionalAddress godoc
// @Summary Get a professional address by ID
// @Description Get a professional address by ID
// @Tags professionalAddresses
// @Accept  json
// @Produce  json
// @Param id path string true "Professional Address ID"
// @Success 200 {object} models.ProfessionalAddress
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /professionalAddresses/{id} [get]
func GetProfessionalAddress(c *fiber.Ctx) error {
	collection := database.GetCollection("professionalAddresses")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, _ := primitive.ObjectIDFromHex(c.Params("id"))
	var address models.ProfessionalAddress
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&address)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(map[string]string{"error": "Professional address not found"})
	}

	return c.Status(http.StatusOK).JSON(address)
}

// UpdateProfessionalAddress godoc
// @Summary Update a professional address
// @Description Update a professional address by ID
// @Tags professionalAddresses
// @Accept  json
// @Produce  json
// @Param id path string true "Professional Address ID"
// @Param address body models.ProfessionalAddress true "Professional Address Payload"
// @Success 200 {object} models.ProfessionalAddress
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /professionalAddresses/{id} [put]
func UpdateProfessionalAddress(c *fiber.Ctx) error {
	collection := database.GetCollection("professionalAddresses")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, _ := primitive.ObjectIDFromHex(c.Params("id"))
	var address models.ProfessionalAddress
	if err := c.BodyParser(&address); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}

	update := bson.M{
		"$set": address,
	}

	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(address)
}

// DeleteProfessionalAddress godoc
// @Summary Delete a professional address
// @Description Delete a professional address by ID
// @Tags professionalAddresses
// @Accept  json
// @Produce  json
// @Param id path string true "Professional Address ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /professionalAddresses/{id} [delete]
func DeleteProfessionalAddress(c *fiber.Ctx) error {
	collection := database.GetCollection("professionalAddresses")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, _ := primitive.ObjectIDFromHex(c.Params("id"))

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(map[string]string{"message": "Professional address deleted"})
}
