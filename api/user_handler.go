package api

import "github.com/gofiber/fiber/v2"

func HandleGetUsers(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"john": "John Wick", "helen": "Helen Wick", "winston": "Winston Scott"})
}