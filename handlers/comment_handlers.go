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

// CreateComment godoc
// @Summary Create a new comment
// @Description Create a new comment on a post
// @Tags comments
// @Accept  json
// @Produce  json
// @Param comment body models.Comment true "Comment Payload"
// @Success 200 {object} models.Comment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /comments [post]
func CreateComment(c *fiber.Ctx) error {
	collection := database.GetCollection("comments")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var comment models.Comment
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}

	comment.CommentID = int(primitive.NewObjectID().Timestamp().Unix())
	comment.CreationDateTime = time.Now()

	_, err := collection.InsertOne(ctx, comment)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(comment)
}

// GetComment godoc
// @Summary Get a comment by ID
// @Description Get a comment by ID
// @Tags comments
// @Accept  json
// @Produce  json
// @Param id path string true "Comment ID"
// @Success 200 {object} models.Comment
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /comments/{id} [get]
func GetComment(c *fiber.Ctx) error {
	collection := database.GetCollection("comments")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, _ := primitive.ObjectIDFromHex(c.Params("id"))
	var comment models.Comment
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&comment)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(map[string]string{"error": "Comment not found"})
	}

	return c.Status(http.StatusOK).JSON(comment)
}

// UpdateComment godoc
// @Summary Update a comment
// @Description Update a comment by ID
// @Tags comments
// @Accept  json
// @Produce  json
// @Param id path string true "Comment ID"
// @Param comment body models.Comment true "Comment Payload"
// @Success 200 {object} models.Comment
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /comments/{id} [put]
func UpdateComment(c *fiber.Ctx) error {
	collection := database.GetCollection("comments")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, _ := primitive.ObjectIDFromHex(c.Params("id"))
	var comment models.Comment
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}

	update := bson.M{
		"$set": comment,
	}

	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(comment)
}

// DeleteComment godoc
// @Summary Delete a comment
// @Description Delete a comment by ID
// @Tags comments
// @Accept  json
// @Produce  json
// @Param id path string true "Comment ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /comments/{id} [delete]
func DeleteComment(c *fiber.Ctx) error {
	collection := database.GetCollection("comments")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, _ := primitive.ObjectIDFromHex(c.Params("id"))

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(map[string]string{"message": "Comment deleted"})
}
