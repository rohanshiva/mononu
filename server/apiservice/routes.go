package apiservice

import (
	"mononu/apiservice/handlers"
	"github.com/gofiber/fiber/v2"
)

func ApiServiceRouter(app fiber.Router)  {
	app.Get("/hello", handlers.GetHello())
}