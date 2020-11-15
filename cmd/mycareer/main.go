package main

import (
	"ephelsa/my-career/pkg/env"
	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig := env.Setup()

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

	app.Listen(envConfig.Server.Port)
}
