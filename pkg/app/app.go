package app

import (
	"ephelsa/my-career/pkg/env"
	"github.com/gofiber/fiber/v2"
	"log"
)

func StartApi(config env.Server) {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
	})

	log.Fatal(app.Listen(config.Port))
}
