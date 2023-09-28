package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadzhuhry/singkating-api-services/config"
	"github.com/muhammadzhuhry/singkating-api-services/helper"
)

func main() {
	app := fiber.New()

	helper.Log.Info("Server started on port: " + config.Port)
	app.Listen(":" + config.Port)
}
