package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Second World yang manis dan cantik!")
	})

	app.Get("/:name", func(c *fiber.Ctx) error {
		str := fmt.Sprintf(`Hello, %v`, c.Params("name"))
		return c.SendString(str)
	})

	app.Listen(":3000")
}
