package handlers

import (
	"context"

	"github.com/Sanjaiy/Library-go/rest-api/database"
	"github.com/Sanjaiy/Library-go/rest-api/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Data transmission object
type libraryDTO struct {
	Name    string   `json:"name" bson:"name"`
	Address string   `json:"address" bson:"address"`
	Empty   []string `json:"empty" bson:"books"`
}

func CreateLibrary(c *fiber.Ctx) error {
	newLibrary := new(libraryDTO)

	if err := c.BodyParser(newLibrary); err != nil {
		return err
	}

	newLibrary.Empty = make([]string, 0)

	libraryCollection := database.GetCollection("libraries")
	nDoc, err := libraryCollection.InsertOne(context.TODO(), newLibrary)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"id": nDoc.InsertedID,
	})
}

func GetLibraries(c *fiber.Ctx) error {
	libraryCollection := database.GetCollection("libraries")
	cursor, err := libraryCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return err
	}

	var libraries []models.Library

	if err = cursor.All(context.TODO(), &libraries); err != nil {
		return err
	}

	return c.JSON(libraries)
}

func DeleteLibrary(c *fiber.Ctx) error {
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

	libraryCollection := database.GetCollection("libraries")

	_, err = libraryCollection.DeleteOne(context.TODO(), bson.M{"_id": libraryObjectID})
	if err != nil {
		return err
	}

	return c.SendString("Library deleted successfully")
}
