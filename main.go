package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/muhammadzhuhry/singkatin-api-services/config"
	"github.com/muhammadzhuhry/singkatin-api-services/db"
	"github.com/muhammadzhuhry/singkatin-api-services/handlers"
	"github.com/muhammadzhuhry/singkatin-api-services/helper"
	"github.com/muhammadzhuhry/singkatin-api-services/repository"
	"github.com/muhammadzhuhry/singkatin-api-services/routes"
	"github.com/muhammadzhuhry/singkatin-api-services/service"
)

func main() {
	app := fiber.New()

	// cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST",
	}))

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

	helper.Log.Info("Server started on port: "+config.Port, config.BasicAuthUsername, config.BasicAuthPassword)
	app.Listen(":" + config.Port)
}
