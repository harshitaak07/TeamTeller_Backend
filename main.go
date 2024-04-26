package main

import (
	"log"
	"ttbackend/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectToDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":8000"))
}
