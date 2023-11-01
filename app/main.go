package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/viilis/go-api/app/db"
	"github.com/viilis/go-api/app/todo"
	"github.com/viilis/go-api/app/utils"
)

func main() {
	log.Println("Staring server")
	app := fiber.New()
	
	// config
	utils.InitConfig()
	
	// database
	db.InitDb()

	// middlewares
	app.Use(helmet.New())
	app.Use(cors.New())	

	todo.TodoRoutes(app)
	
	// simple 404-handler
	app.Use(func (c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	app.Listen(":"+ utils.Config.Port)
}