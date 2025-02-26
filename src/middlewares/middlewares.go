package middlewares

import (
	"github.com/gofiber/fiber/v3"
)

func RegisterMiddlewares(app *fiber.App) {
	app.Use(AuthMiddleware())
}
