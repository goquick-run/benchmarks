package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/v1/user/:id", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		myuser := struct {
			Name string `json:"name"`
		}{
			Name: c.Params("id"),
		}
		return c.JSON(myuser)
	})

	app.Listen(":8080")
}
