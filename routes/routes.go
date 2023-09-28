package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadzhuhry/singkating-api-services/auth"
	"github.com/muhammadzhuhry/singkating-api-services/handlers"
)

func SetupURLRoute(app *fiber.App, urlHandler *handlers.UrlHandler) {
	v1 := app.Group("/api/v1")
	v1.Get("/", auth.BasicAuthMiddleware(), urlHandler.Root)
}
