package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadzhuhry/singkating-api-services/helper"
	"github.com/muhammadzhuhry/singkating-api-services/models"
	"github.com/muhammadzhuhry/singkating-api-services/models/request"
	"github.com/muhammadzhuhry/singkating-api-services/service"
)

type UrlHandler struct {
	UrlService service.UrlService
	Validate   *validator.Validate
}

func NewUrlHandler(urlService *service.UrlService, validate *validator.Validate) *UrlHandler {
	return &UrlHandler{
		UrlService: *urlService,
		Validate:   validate,
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

	response := h.UrlService.UrlShorten(c, payload)
	return helper.SendResponse(c, response)
}

func (h *UrlHandler) RedirectUrl(c *fiber.Ctx) error {
	url := c.Query("url")

	err := h.Validate.Var(url, "required")
	if err != nil {
		response := helper.ErrorResponse(fiber.StatusBadRequest, nil, err.Error())
		return helper.SendResponse(c, response)
	}

	response := h.UrlService.RedirectUrl(c, url)
	return helper.SendResponse(c, response)
}
