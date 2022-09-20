package main

import (
	"github.com/dogab/finance-analyser-api/pkg/config"
	"github.com/dogab/finance-analyser-api/pkg/database"
	"github.com/dogab/finance-analyser-api/pkg/router"
	"github.com/gofiber/fiber/v2"
)

// @title Financeanalyser API
// @version 0.1.0
// @description Swagger endpoint
// @termsOfService http://swagger.io/terms/
// @contact.name Dominic Gabriel
// @contact.email github.com/dogab
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:5000
// @BasePath /
func main() {
	// Init envs
	config.InitEnvs()

	// Start a new fiber app
	app := fiber.New()

	// Connect to the Database
	database.ConnectDB()

	// setup the router
	router.SetupRoutes(app)

	// Listen on PORT 300
	if err := app.Listen(":5000"); err != nil {
		panic(err)
	}
}
