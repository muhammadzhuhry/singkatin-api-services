package service

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/muhammadzhuhry/singkating-api-services/helper"
	"github.com/muhammadzhuhry/singkating-api-services/models"
	"github.com/muhammadzhuhry/singkating-api-services/models/entity"
	"github.com/muhammadzhuhry/singkating-api-services/models/request"
	"github.com/muhammadzhuhry/singkating-api-services/repository"
)

type UrlService struct {
	UrlRepository repository.UrlRepository
}

func NewUrlService(urlRepository *repository.UrlRepository) *UrlService {
	return &UrlService{
		UrlRepository: *urlRepository,
	}
}

func (s *UrlService) UrlShorten(c *fiber.Ctx, request request.Url) models.Response {
	payload := entity.Url{
		LongUrl:   request.LongUrl,
		ShortUrl:  helper.GenerateShortUrl(),
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(time.Hour * 24),
	}

	// logic check duplicate short url and long url if exist return it

	url, err := s.UrlRepository.InsertUrl(&payload)
	if err != nil {
		response := helper.ErrorResponse(fiber.StatusInternalServerError, nil, "Failed to shortning url")
		return response
	}

	response := helper.SuccessResponse(fiber.StatusCreated, payload, "Success shortning url: "+url.LongUrl)
	return response
}
