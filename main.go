package main

import (
	"github.com/Arenelin/List-of-current-affairs/src/middlewares"
	"github.com/Arenelin/List-of-current-affairs/src/models"
	"github.com/Arenelin/List-of-current-affairs/src/routes"
	"github.com/Arenelin/List-of-current-affairs/src/utils"
	"log"
)

func main() {
	utils.LoadEnv()
	models.OpenDatabaseConnection()
	models.AutoMigrateModels()
	router := routes.SetupRoutes()
	middlewares.RegisterMiddlewares(router)

	log.Fatal(router.Listen(":3000"))
}
