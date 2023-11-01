package todo

import "github.com/gofiber/fiber/v2"

func TodoRoutes(app fiber.Router) {
	router := app.Group("/todo")

	//POST
	router.Post("/create", PostTodoHandler)

	// GET
	router.Get("/all", GetTodosHandler)
	router.Get("/:id", GetOneTodoHandler)

	//PUT
	router.Patch("/:id", UpdateTodoHandler)

	//DEL
	router.Delete("/:id", DeleteTodoHandler)
}