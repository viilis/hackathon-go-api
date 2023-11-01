package todo

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/viilis/go-api/app/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllTodos() ([]todo, error) {
	var todoResults []todo

	cursor, err := db.DbRef.Collection("todos").Find(context.TODO(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	err = cursor.All(context.TODO(), &todoResults)

	if err != nil {
		log.Fatal(err)
	}

	return todoResults, nil
}

func FindOneTodo(objId primitive.ObjectID) (todo, error) {
	var result todo
	
	filter := bson.M{ "_id": objId }

	err := db.DbRef.Collection("todos").FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	return result, err
}

func CreateNewTodo(body postTodoRequest) error {

	todo := todo {
		Title: body.Title,
		Completed: false,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if _, err := db.DbRef.Collection("todos").InsertOne(context.TODO(), todo); err != nil {
		log.Fatal(err)
	}

	return nil
}

func DeleteTodo(id string) error  {

	hexToId, err := primitive.ObjectIDFromHex(id)

	//TODO: Move conversion to handler and return right statuscode if cannot be converted (=invalid id)
	if err != nil {
		return errors.New("could not convert id")
	}
	
	filter := bson.M{ "_id": hexToId }

	if _, err := db.DbRef.Collection("todos").DeleteOne(context.TODO(), filter); err != nil {
		log.Fatal(err)
	}

	return nil
}

func UpdateTodo(objId primitive.ObjectID, updates putTodoRequest) (todo, error) {
	
	oldTodo, err := FindOneTodo(objId)

	if err != nil {
		log.Fatal(err)
	}

	updatedTodo := todo {
		Id: oldTodo.Id,
		Title: updates.Title,
		Completed: updates.Completed,
		CreatedAt: oldTodo.CreatedAt,
		UpdatedAt: time.Now().Unix(),
	}

	filter := bson.M{ "_id": objId }

	if _, err := db.DbRef.Collection("todos").ReplaceOne(context.TODO(), filter, updatedTodo); err != nil {
		log.Fatal(err)
	}

	return updatedTodo, err
}