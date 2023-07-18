package models

// Uppercase Public - lowercase private
type Song struct {
	ID     string  `json:"id" bson:"_id"`
	Title  string  `json:"title" bson:"title"`
	Author string  `json:"author" bson:"author"`
	Length float64 `json:"length" bson:"length"`
}
