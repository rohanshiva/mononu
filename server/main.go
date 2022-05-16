package main

import (
	"log"
	"mononu/apiservice"
	"mononu/staticservice"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from mononu server (trying to render multiple frontend codebases) 👋🏿")
	})

	// sample api service
	api := app.Group("/api")
	apiservice.ApiServiceRouter(api)


	// rendering static site
	app.Static("/static", "../staticapp")
	staticapp := app.Group("/static")
	staticservice.StaticServiceRouter(staticapp)

	// rendering pre built react app
	app.Static("/reactapp", "../reactapp/build")
	app.Static("*", "../reactapp/build/index.html")


	log.Fatal(app.Listen("127.0.0.1:3000"))
}