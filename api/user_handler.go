package api

import (
	"errors"

	"github.com/RogerWaldron/go-reserveration-api/db"
	"github.com/RogerWaldron/go-reserveration-api/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler( userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(users)
}

func (h *UserHandler) HandleGetUserByID(c *fiber.Ctx) error {
	var (
		id 	= c.Params("id")
	)

	user, err := h.userStore.GetUserByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ErrNotFound(id)
		}
		return err
	}
	
	return c.JSON(user)
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	err := c.BodyParser(&params)
	if err != nil {
		return err
	}

	errors := params.Validate()
	if len(errors) > 0 {
		return c.JSON(errors)
	}

	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}

	insertedUser, err := h.userStore.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}

	return c.JSON(insertedUser)
}

func (h *UserHandler) HandlePutUser(c *fiber.Ctx) error {
	var (
		userID = c.Params("id")
		params types.UpdateUserParams
	)
	err := c.BodyParser(&params)
	if err != nil {
		return err
	}

	filter := db.Map{"_id": userID}
	err = h.userStore.UpdateUser(c.Context(), filter, params)
	if err != nil {
		return err
	}

	return c.JSON(map[string]string{"updated": userID})
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	deleteID := c.Params("id")
	if (deleteID == "") {
		return ErrInvalidID(deleteID)
	}
	err := h.userStore.DeleteUser(c.Context(), deleteID)
	if err != nil {
		return err
	}

	return c.JSON(map[string]string{"Deleted": deleteID})
}