package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadzhuhry/singkatin-api-services/auth"
	"github.com/muhammadzhuhry/singkatin-api-services/handlers"
)

func SetupURLRoute(app *fiber.App, urlHandler *handlers.UrlHandler) {
	v1 := app.Group("/api/v1/url")
	v1.Get("/", auth.BasicAuthMiddleware(), urlHandler.Root)
	v1.Post("/shorten", auth.BasicAuthMiddleware(), urlHandler.UrlShorten)
	v1.Get("/redirect", auth.BasicAuthMiddleware(), urlHandler.RedirectUrl)
}
