package main

import (
	"flag"

	"github.com/RogerWaldron/go-reserveration-api/api"
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
	appv1.Get("/users", api.HandleGetUsers)
	app.Listen(*listenAddress) 
}