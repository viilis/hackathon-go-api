package todo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type todo struct {
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty" validate:"required"`
	Title string `json:"title" bson:"title" validate:"required,min=1,max=255"`
	Completed bool `json:"completed" bson:"completed" validate:"required"`
	CreatedAt int64 `json:"createdAt" bson:"createdAt" validate:"required"`
	UpdatedAt int64 `json:"updatedAt" bson:"updatedAt" validate:"required"`
}

type postTodoRequest struct {
	Title string `json:"title" validate:"required,min=1,max=255"`
}

type putTodoRequest struct {
	Title string `json:"title" bson:"title,omitempty" validate:"min=1,max=255"`
	Completed bool `json:"completed" bson:"completed,omitempty"`
}