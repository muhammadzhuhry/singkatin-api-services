package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadzhuhry/singkating-api-services/config"
	"github.com/muhammadzhuhry/singkating-api-services/handlers"
	"github.com/muhammadzhuhry/singkating-api-services/helper"
	"github.com/muhammadzhuhry/singkating-api-services/routes"
)

func main() {
	app := fiber.New()

	// initialize handlers for url
	urlHandler := handlers.NewUrlHandler()

	// setup routes
	routes.SetupURLRoute(app, urlHandler)

	helper.Log.Info("Server started on port: " + config.Port)
	app.Listen(":" + config.Port)
}
