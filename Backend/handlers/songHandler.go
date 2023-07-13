package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type newSongDTO struct {
	Title    string  `json:"title" bson:"title" validate:"required"`
	Author   string  `json:"author" bson:"author" validate:"required"`
	Length   float64 `json:"length" bson:"length" validate:"required"`
	LibaryId string  `json:"libaryId" bson:"libaryId" validate:"required"`
}

func CreateSong(c *fiber.Ctx) error {
	nSong := new(newSongDTO)
	if err := c.BodyParser(nSong); err != nil {
		return err
	}
	fmt.Println(nSong)
	return nil
}
