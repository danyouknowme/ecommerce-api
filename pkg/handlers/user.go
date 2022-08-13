package handlers

import (
	"log"

	"github.com/danyouknowme/ecommerce/pkg/database/dbmodels"
	"github.com/gofiber/fiber/v2"
)

func RegisterAPI(ctx *fiber.Ctx) error {
	var newUser dbmodels.User
	if err := ctx.BodyParser(&newUser); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	log.Printf("post: /api/v1/users/register")

	err := dbmodels.Register(newUser)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Create a new user successfully!",
	})
}

func LoginAPI(ctx *fiber.Ctx) error {
	var req dbmodels.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	log.Printf("post: /api/v1/users/login")
	user, err := dbmodels.Login(req.Username, req.Password)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}
