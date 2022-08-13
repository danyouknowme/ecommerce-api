package handlers

import (
	"log"

	"github.com/danyouknowme/ecommerce/pkg/database/dbmodels"
	"github.com/gofiber/fiber/v2"
)

type UserResponse struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

func RegisterAPI(ctx *fiber.Ctx) error {
	var req dbmodels.User
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	log.Printf("post: /api/v1/users/register")

	err := dbmodels.Register(req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	userResponse := UserResponse{
		Username: req.Username,
		FullName: req.FullName,
		Email:    req.Email,
	}

	return ctx.JSON(fiber.Map{
		"message": "Create a new user successfully!",
		"user":    userResponse,
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

func GetUserAPI(ctx *fiber.Ctx) error {
	username := ctx.Params("username")

	log.Printf("get: /api/auth/v1/users/%s", username)

	user, err := dbmodels.GetUser(username)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	userResponse := UserResponse{
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
	}

	return ctx.Status(fiber.StatusOK).JSON(userResponse)
}
