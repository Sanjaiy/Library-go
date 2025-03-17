package main

import (
	"os"

	"github.com/Sanjaiy/Library-go/rest-api/database"
	"github.com/Sanjaiy/Library-go/rest-api/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := initApp()
	if err != nil {
		panic(err)
	}

	// Close MongoDB
	defer database.CloseMongoDB()

	app := generateApp()


	// get the port from the env
	port := os.Getenv("PORT")

	app.Listen(":" + port)
}

func loadEnv() error {
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
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


func generateApp() *fiber.App {
	app := fiber.New()

	// create health check route
	app.Get("/health", func (c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// create the library groups and routes
	libGroup := app.Group("/library")
	libGroup.Get("/", handlers.GetLibraries)
	libGroup.Post("/", handlers.CreateLibrary)
	libGroup.Post("/:id/book", handlers.CreateBook)
	libGroup.Delete("/:id", handlers.DeleteLibrary)
	return app
}