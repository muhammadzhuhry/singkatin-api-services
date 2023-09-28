package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadzhuhry/singkating-api-services/models"
)

func SuccessResponse(code int, data interface{}, message string) models.Response {
	return models.Response{
		Code:    code,
		Status:  true,
		Data:    data,
		Message: message,
	}
}

func ErrorResponse(code int, data interface{}, message string) models.Response {
	return models.Response{
		Code:    code,
		Status:  false,
		Data:    data,
		Message: message,
	}
}

func SendResponse(c *fiber.Ctx, res models.Response) error {
	return c.Status(res.Code).JSON(res)
}
