package db

import "time"

type Todo struct {
	Id int `bson:"_id"`
	Title string
	Completed bool
	CreatedAt time.Time
}