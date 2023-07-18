package models

type Libary struct {
	ID    string `json:"id" bson:"_id"`
	Name  string `json:"name" bson:"name"`
	Album string `json:"album" bson:"album"`
	Songs []Song `json:"songs" bson:"songs"`
}
