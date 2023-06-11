package api

import (
	"github.com/RogerWaldron/go-reserveration-api/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"john": "John Wick", "helen": "Helen Wick", "winston": "Winston Scott"})
}

func HandleGetUserByID(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "John",
		LastName: "Wick",
	}
	return c.JSON(u)
}