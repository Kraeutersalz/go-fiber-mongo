package handlers

import "github.com/gofiber/fiber/v2"

func FirstHandler(c *fiber.Ctx) error {
	return c.SendString("It works! in FirstHandler/Libary")
}
