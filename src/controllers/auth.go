package controllers

import (
	"fmt"
	"github.com/Arenelin/List-of-current-affairs/src/models"
	"github.com/gofiber/fiber/v3"
)

func RegisterUser(ctx fiber.Ctx) error {
	var newUser *models.User
	fmt.Println(newUser)
	if err := ctx.Bind().JSON(&newUser); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User should have a email and password",
		})
	}

	user, err := newUser.Register()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to register user",
		})

	}
	ctx.Status(fiber.StatusCreated).JSON(user)
	return nil
}

func LoginUser(ctx fiber.Ctx) error {
	var user models.User
	fmt.Println(user)
	if err := ctx.Bind().JSON(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Incorrect login or password",
		})

	}
	authorizedUser, err := user.Login()

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Incorrect login or password",
		})

	}
	ctx.Status(fiber.StatusCreated).JSON(authorizedUser)
	return nil
}
