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

// CreateForum godoc
// @Summary Create a new forum post
// @Description Create a new forum post by a user
// @Tags forums
// @Accept  json
// @Produce  json
// @Param forum body models.Forum true "Forum Payload"
// @Success 200 {object} models.Forum
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /forums [post]
func CreateForum(c *fiber.Ctx) error {
    collection := database.GetCollection("forums")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var forum models.Forum
    if err := c.BodyParser(&forum); err != nil {
        return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
    }

    forum.ID = primitive.NewObjectID().Hex()
    forum.CreatedAt = time.Now()

    _, err := collection.InsertOne(ctx, forum)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
    }

    return c.Status(http.StatusOK).JSON(forum)
}

// GetForum godoc
// @Summary Get a forum post by ID
// @Description Get a forum post by ID
// @Tags forums
// @Accept  json
// @Produce  json
// @Param id path string true "Forum ID"
// @Success 200 {object} models.Forum
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /forums/{id} [get]
func GetForum(c *fiber.Ctx) error {
    collection := database.GetCollection("forums")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    id, _ := primitive.ObjectIDFromHex(c.Params("id"))
    var forum models.Forum
    err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&forum)
    if err != nil {
        return c.Status(http.StatusNotFound).JSON(map[string]string{"error": "Forum post not found"})
    }

    return c.Status(http.StatusOK).JSON(forum)
}

// Add other handlers for update and delete with similar annotations
