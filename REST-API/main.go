package main

import (
	"os"

	"github.com/Sanjaiy/Library-go/rest-api/database"
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