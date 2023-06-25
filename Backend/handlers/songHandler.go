package handlers

type newSongDTO struct {
	Title    string `json:"title" bson:"title" validate:"required"`
	Author   string `json:"author" bson:"author" validate:"required"`
	Length   int    `json:"length" bson:"length" validate:"required"`
	LibaryId string `json:"libaryId" bson:"libaryId" validate:"required"`
}

func CreateSong() {

}
