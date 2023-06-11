package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	appv1 := app.Group("/api/v1")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
	appv1.Get("/users", handleUser)
	app.Listen(":3000")
}

func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "John Wick"})
}