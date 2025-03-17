package handlers

import (
	"context"

	"github.com/Sanjaiy/Library-go/rest-api/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookDTO struct {
	Title string `json:"title" bson:"title"`
	Author string `json:"author" bson:"author"`
	ISBN string `json:"isbn" bson:"isbn"`
}

func CreateBook(c *fiber.Ctx) error {
	newBook := new(BookDTO)
	
	if err := c.BodyParser(newBook); err != nil {
		return err
	}

	libraryId := c.Params("id")
	if libraryId == "" {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": "Library ID is required",
		})
	}

	libraryObjectID, err := primitive.ObjectIDFromHex(libraryId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Library ID",
		})
	}

	collection := database.GetCollection("libraries")

	filter := bson.M{"_id": libraryObjectID}
	updatePayload := bson.M{"$push": bson.M{"books": newBook}}


	_, err = collection.UpdateOne(context.TODO(), filter, updatePayload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	
	return c.SendString("Book created successfully")
}