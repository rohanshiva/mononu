package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetHello() fiber.Handler {
	return func (c *fiber.Ctx) error {
		return c.SendString("hello from sample api service")
	}
}