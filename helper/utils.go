package helper

import (
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ParseRequestBody(c *fiber.Ctx, payload interface{}) error {
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	return nil
}

func GenerateShortUrl() string {
	var length = 7
	const characterSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.Seed(time.Now().UnixNano())

	var shortURL string

	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(characterSet))
		randomChar := characterSet[randomIndex]

		shortURL += string(randomChar)
	}

	return shortURL
}
