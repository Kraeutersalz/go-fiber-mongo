package models

type Libary struct {
	ID     string `json:"id" bson:"_id"`
	Name   string `json:"name" bson:"name"`
	Adress string `json:"adress" bson:"adress"`
	Songs  []Song `json:"songs" bson:"songs"`
}
