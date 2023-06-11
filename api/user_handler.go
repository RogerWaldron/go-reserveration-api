package api

import "github.com/gofiber/fiber/v2"

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"john": "John Wick", "helen": "Helen Wick", "winston": "Winston Scott"})
}