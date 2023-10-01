package auth

import (
	"encoding/base64"

	"github.com/gofiber/fiber/v2"
	"github.com/muhammadzhuhry/singkatin-api-services/config"
	"github.com/muhammadzhuhry/singkatin-api-services/helper"
)

func BasicAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := config.BasicAuthUsername
		password := config.BasicAuthPassword

		authString := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		authHeader := "Basic " + authString

		if c.Get("Authorization") != authHeader {
			res := helper.ErrorResponse(fiber.StatusUnauthorized, nil, "Unauthorized")
			return helper.SendResponse(c, res)
		}

		return c.Next()
	}
}
