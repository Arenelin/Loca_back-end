package main

import (
	"fmt"
	"github.com/Arenelin/List-of-current-affairs/src/models"
	"github.com/Arenelin/List-of-current-affairs/src/utils"
	t "github.com/Arenelin/List-of-current-affairs/types"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	utils.LoadEnv()
	app := fiber.New()

	app.Post("/login", func(ctx fiber.Ctx) error {
		loginData := new(t.AuthDTO)

		if err := ctx.Bind().JSON(&loginData); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "User should have a email and password",
			})
		}
		fmt.Println(loginData)
		return ctx.Status(fiber.StatusOK).JSON(loginData)
	})

	app.Get("/profile", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	models.OpenDatabaseConnection()

	log.Fatal(app.Listen(":3000"))
}
