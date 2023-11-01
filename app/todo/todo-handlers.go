package todo

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTodosHandler(c *fiber.Ctx) error {
	res, err := FindAllTodos()
	
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(res)
}

func GetOneTodoHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	if len(id) == 0 {
		return c.Status(400).SendString("No id")
	}

	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(500).SendString("Failed to convert id")
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
		return c.SendStatus(500)
	}

	if err := CreateNewTodo(body); err != nil {
		return c.SendStatus(500)
	}

	return c.SendStatus(200)
}

func UpdateTodoHandler(c *fiber.Ctx) error {
	var body putTodoRequest

	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(500)
	}

	id := c.Params("id")

	if len(id) == 0 {
		return c.Status(400).SendString("No id")
	}

	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(500).SendString("Failed to convert id")
	}

	res, err := UpdateTodo(objId, body)

	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(res)
}

func DeleteTodoHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	if len(id) == 0 {
		return c.Status(400).SendString("No id")
	}

	if err := DeleteTodo(id); err != nil {
		return c.SendStatus(500)
	}

	return c.SendStatus(200)
}
