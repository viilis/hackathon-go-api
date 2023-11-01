package todo

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

func GetTodosHandler(c *fiber.Ctx) error {
	res, err := FindAllTodos()
	
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(res)
}

func GetOneTodoHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	err := validate.Var(id, "required,mongodb")

	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "Validation error")
	}

	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	res, err := FindOneTodo(objId)

	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(res)
}

func PostTodoHandler(c *fiber.Ctx) error {
	var body postTodoRequest

	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}

	err := validate.Struct(body)

	if err != nil {
		fmt.Println(err)
		return fiber.NewError(fiber.ErrBadRequest.Code, "Validation error")
	}

	if err := CreateNewTodo(body); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(200)
}

func UpdateTodoHandler(c *fiber.Ctx) error {
	var body putTodoRequest

	id := c.Params("id")
	err := validate.Var(id, "required,mongodb")

	//TODO: Get rid of repetitice error handling -> create validator-bodyparses middleware
	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "Validation error")
	}

	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}

	err = validate.Struct(body)

	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "Validation error")
	}

	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	res, err := UpdateTodo(objId, body)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(res)
}

func DeleteTodoHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	err := validate.Var(id, "required,mongodb")

	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "Validation error")
	}

	if err := DeleteTodo(id); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(200)
}
