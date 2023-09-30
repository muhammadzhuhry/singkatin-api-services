package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadzhuhry/singkating-api-services/config"
	"github.com/muhammadzhuhry/singkating-api-services/db"
	"github.com/muhammadzhuhry/singkating-api-services/handlers"
	"github.com/muhammadzhuhry/singkating-api-services/helper"
	"github.com/muhammadzhuhry/singkating-api-services/repository"
	"github.com/muhammadzhuhry/singkating-api-services/routes"
	"github.com/muhammadzhuhry/singkating-api-services/service"
)

func main() {
	app := fiber.New()

	// initialize new db
	db := db.NewDB()
	defer db.Close()

	// initialize validator
	validate := validator.New()

	// initialize handlers for url
	urlRepository := repository.NewUrlRepository(db)
	urlService := service.NewUrlService(urlRepository)
	urlHandler := handlers.NewUrlHandler(urlService, validate)

	// setup routes
	routes.SetupURLRoute(app, urlHandler)

	helper.Log.Info("Server started on port: " + config.Port)
	app.Listen(":" + config.Port)
}
