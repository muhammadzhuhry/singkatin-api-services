package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadzhuhry/singkating-api-services/helper"
	"github.com/muhammadzhuhry/singkating-api-services/models"
)

type UrlHandler struct{}

func NewUrlHandler() *UrlHandler {
	return &UrlHandler{}
}

func (h *UrlHandler) Root(c *fiber.Ctx) error {
	return helper.SendResponse(c, models.Response{})
}
