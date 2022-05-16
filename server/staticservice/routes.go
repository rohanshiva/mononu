package staticservice

import (
	"github.com/gofiber/fiber/v2"
)

func StaticServiceRouter(app fiber.Router) {
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello from static service")
	})
}