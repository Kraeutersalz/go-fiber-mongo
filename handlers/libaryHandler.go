package handlers

import (
	"context"

	"github.com/Kraeutersalz/go-fiber-mongo/database"
	"github.com/Kraeutersalz/go-fiber-mongo/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type libaryDTO struct {
	Name  string `json:"name" bson:"name" validate:"required"`
	Album string `json:"album" bson:"album" validate:"required"`
}

// GET
func GetLibaries(c *fiber.Ctx) error {
	libaryCollection := database.GetCollection("libaries")
	cursor, err := libaryCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		return err
	}
	var libaries []models.Libary
	if err := cursor.All(context.TODO(), &libaries); err != nil {
		return err
	}
	return c.JSON(libaries)
}

// POST
func CreateLibary(c *fiber.Ctx) error {
	nLibary := new(libaryDTO)
	if err := c.BodyParser(nLibary); err != nil {
		return err
	}

	libaryCollection := database.GetCollection("libaries")
	insertOneResult, err := libaryCollection.InsertOne(context.TODO(), nLibary)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"id": insertOneResult.InsertedID})
}
