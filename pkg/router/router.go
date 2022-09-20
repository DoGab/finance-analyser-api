package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/dogab/finance-analyser-api/pkg/routes"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	app.Use(logger.New())
	app.Use(cors.New())
	routes.SetupSwaggerRoutes(app)

	// Group api calls with param "/api"
	// api := app.Group("/api")
	// api.Use(middleware.JWTProtected())

}
