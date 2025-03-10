package main

import (
	"context"

	"github.com/Sanjaiy/Library-go/rest-api/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	err := initApp()
	if err != nil {
		panic(err)
	}

	// Close MongoDB
	defer database.CloseMongoDB()

	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		// write a todo to the database
		sampleDoc := bson.M{"name": "sample todo"}
		collection := database.GetCollection("todos")
		nDoc, err := collection.InsertOne(context.TODO(), sampleDoc)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error inserting todo")
		}

		return c.JSON(nDoc)
	})

	app.Listen(":3000")
}

func loadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	return nil
}


func initApp() error {
	// Setup ENV
	err := loadEnv()
	if err != nil {
		return err
	}

	// Start MongoDB
	err = database.StartMongoDB()
	if err != nil {
		return err
	}

	return nil
}