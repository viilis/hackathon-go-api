package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/viilis/go-api/cmd/todo"
)

func main() {
	app := fiber.New()

	// middlewares
	app.Use(helmet.New())
	app.Use(cors.New())
	
	api := app.Group("/api")

	api.Get("/", todo.TodoHandler)

	// simple 404-handler
	app.Use(func (c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	app.Listen(":3000")
}