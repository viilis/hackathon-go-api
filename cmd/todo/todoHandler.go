package todo

import (
	"github.com/gofiber/fiber/v2"
)

func TodoHandler(c *fiber.Ctx) error {
	return c.SendStatus(200)
}