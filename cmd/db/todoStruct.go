package db

type Todo struct {
	Id int `bson:"_id"`
	Title string
}