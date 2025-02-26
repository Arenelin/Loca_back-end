package routes

import (
	"github.com/Arenelin/List-of-current-affairs/src/controllers"
	"github.com/gofiber/fiber/v3"
)

func authGroupRouter(app *fiber.App) {
	auth := app.Group("/api/v1/auth")
	auth.Post("/register", controllers.RegisterUser)
	auth.Post("/login", controllers.LoginUser)
}

func SetupRoutes() *fiber.App {
	app := fiber.New()
	authGroupRouter(app)
	return app
}
