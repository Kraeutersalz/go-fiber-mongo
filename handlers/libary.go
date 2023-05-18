package handlers

import (
	"context"

	"github.com/Kraeutersalz/go-fiber-mongo/database"
	"github.com/gofiber/fiber/v2"
)

type libaryDTO struct {
	Name  string `json:"name" bson:"name" validate:"required"`
	Album string `json:"album" bson:"album" validate:"required"`
}

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
