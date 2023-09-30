package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadzhuhry/singkating-api-services/helper"
	"github.com/muhammadzhuhry/singkating-api-services/models"
	"github.com/muhammadzhuhry/singkating-api-services/models/request"
)

type UrlHandler struct {
	Validate *validator.Validate
}

func NewUrlHandler(validate *validator.Validate) *UrlHandler {
	return &UrlHandler{
		Validate: validate,
	}
}

func (h *UrlHandler) Root(c *fiber.Ctx) error {
	return helper.SendResponse(c, models.Response{})
}

func (h *UrlHandler) UrlShorten(c *fiber.Ctx) error {
	var payload request.Url

	if err := helper.ParseRequestBody(c, &payload); err != nil {
		response := helper.ErrorResponse(fiber.StatusBadRequest, nil, "Invalid request")
		return helper.SendResponse(c, response)
	}

	err := h.Validate.Struct(payload)
	if err != nil {
		response := helper.ErrorResponse(fiber.StatusBadRequest, nil, err.Error())
		return helper.SendResponse(c, response)
	}

	response := models.Response{
		Code:    200,
		Status:  true,
		Data:    nil,
		Message: "Success shortning url",
	}
	return helper.SendResponse(c, response)
}
