package main

import (
	"github.com/Kraeutersalz/go-fiber-mongo/handlers"
	"github.com/gofiber/fiber/v2"
)

func generateApp() *fiber.App {
	app := fiber.New()

	//Create health check
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	//Create libary group
	libGroup := app.Group("/libary")
	//Routes from libary group
	libGroup.Get("/", handlers.GetLibaries)
	libGroup.Post("/", handlers.CreateLibary)
	libGroup.Post("/song", handlers.CreateSong)

	return app
}
