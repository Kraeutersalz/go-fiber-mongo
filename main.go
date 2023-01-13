package main

import (
	"context"

	"github.com/Kraeutersalz/go-fiber-mongo/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	//init app
	err := initApp()
	if err != nil {
		panic(err)
	}

	//defer close database
	defer database.CloseMongoDB()

	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		//write a title to database
		sampleSong := bson.M{"name": "Sample Song"}
		collection := database.GetCollection("Songs")
		nDoc, err := collection.InsertOne(context.TODO(), sampleSong)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error inserting Sond")
		}

		//send down info about song
		return c.JSON(nDoc)
	})

	app.Listen(":3000")
}

func initApp() error {
	//setup ENV
	err := LoadENV()
	if err != nil {
		return err
	}

	//setup Database
	err = database.StartMongoDB()
	if err != nil {
		return err
	}
	return nil
}

func LoadENV() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
