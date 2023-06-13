package api

import (
	"context"

	"github.com/RogerWaldron/go-reserveration-api/db"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"john": "John Wick", "helen": "Helen Wick", "winston": "Winston Scott"})
}

func (h *UserHandler) HandleGetUserByID(c *fiber.Ctx) error {
	var (
		id 	= c.Params("id")
		ctx = context.Background()
	)

	user, err := h.userStore.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	
	return c.JSON(user)
}