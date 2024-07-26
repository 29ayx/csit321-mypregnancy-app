package handlers

import (
	"context"
	"gofiber-mongodb/models"
	"gofiber-mongodb/server/database"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateConsultationNote godoc
// @Summary Create a new consultation note
// @Description Create a new consultation note for a request
// @Tags consultationnotes
// @Accept  json
// @Produce  json
// @Param note body models.ConsultationNotes true "Consultation Note Payload"
// @Success 200 {object} models.ConsultationNotes
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /consultationnotes [post]
func CreateConsultationNote(c *fiber.Ctx) error {
	collection := database.GetCollection("consultationnotes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var note models.ConsultationNotes
	if err := c.BodyParser(&note); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}

	_, err := collection.InsertOne(ctx, note)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(note)
}

// GetConsultationNote godoc
// @Summary Get a consultation note by Request ID
// @Description Get a consultation note by Request ID
// @Tags consultationnotes
// @Accept  json
// @Produce  json
// @Param id path int true "Request ID"
// @Success 200 {object} models.ConsultationNotes
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /consultationnotes/{id} [get]
func GetConsultationNote(c *fiber.Ctx) error {
	collection := database.GetCollection("consultationnotes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	var note models.ConsultationNotes
	err := collection.FindOne(ctx, bson.M{"requestID": id}).Decode(&note)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(map[string]string{"error": "Consultation note not found"})
	}

	return c.Status(http.StatusOK).JSON(note)
}

// UpdateConsultationNote godoc
// @Summary Update a consultation note
// @Description Update a consultation note by Request ID
// @Tags consultationnotes
// @Accept  json
// @Produce  json
// @Param id path int true "Request ID"
// @Param note body models.ConsultationNotes true "Consultation Note Payload"
// @Success 200 {object} models.ConsultationNotes
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /consultationnotes/{id} [put]
func UpdateConsultationNote(c *fiber.Ctx) error {
	collection := database.GetCollection("consultationnotes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	var note models.ConsultationNotes
	if err := c.BodyParser(&note); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}

	update := bson.M{
		"$set": note,
	}

	_, err := collection.UpdateOne(ctx, bson.M{"requestID": id}, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(note)
}

// DeleteConsultationNote godoc
// @Summary Delete a consultation note
// @Description Delete a consultation note by Request ID
// @Tags consultationnotes
// @Accept  json
// @Produce  json
// @Param id path int true "Request ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /consultationnotes/{id} [delete]
func DeleteConsultationNote(c *fiber.Ctx) error {
	collection := database.GetCollection("consultationnotes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")

	_, err := collection.DeleteOne(ctx, bson.M{"requestID": id})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(map[string]string{"message": "Consultation note deleted"})
}
