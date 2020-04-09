package main

import (
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.JSON(fiber.Map{
			"ping": "pong",
		})
	})

	app.Listen(3000)
}
