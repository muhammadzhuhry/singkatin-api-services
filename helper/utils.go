package helper

import "github.com/gofiber/fiber/v2"

func ParseRequestBody(c *fiber.Ctx, payload interface{}) error {
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	return nil
}
