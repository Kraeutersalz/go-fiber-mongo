package handlers

import (
	"fmt"

	"github.com/Kraeutersalz/go-fiber-mongo/database"
	"github.com/Kraeutersalz/go-fiber-mongo/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	// get the collection reference
	coll := database.GetCollection("libaries")
	//get the filter
	objID, _ := primitive.ObjectIDFromHex(nSong.LibaryId)
	filter := bson.M{"_id": objID}

	fmt.Println(filter)
	nSongData := models.Song{
		Title:  nSong.Title,
		Author: nSong.Author,
		Length: nSong.Length,
	}
	updatePayload := bson.M{"$push": bson.M{"songs": nSongData}}
	//update the document
	_, err := coll.UpdateOne(c.Context(), filter, updatePayload)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.SendString("Song added to libary")
}
