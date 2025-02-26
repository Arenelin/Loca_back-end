package middlewares

import (
	"github.com/gofiber/fiber/v3"
)

func AuthMiddleware() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		return nil
	}
}
