package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddress := flag.String("listenAddress", ":3200", "This listen address for API server")
	flag.Parse()

	app := fiber.New()
	appv1 := app.Group("/api/v1")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
	appv1.Get("/users", handleUser)
	app.Listen(*listenAddress)
}

func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "John Wick"})
}