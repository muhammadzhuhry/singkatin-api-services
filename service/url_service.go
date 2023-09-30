package service

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/muhammadzhuhry/singkating-api-services/config"
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
	location := time.FixedZone(config.Timezone, 7*60*60)

	payload := entity.Url{
		LongUrl:   request.LongUrl,
		ShortUrl:  helper.GenerateShortUrl(),
		CreatedAt: time.Now().In(location).Format("2006-01-02 15:04:05"),
		ExpiredAt: time.Now().In(location).Add(time.Hour * 24).Format("2006-01-02 15:04:05"),
	}

	// logic check duplicate url
	data, _ := s.UrlRepository.FindExistedUrl(payload.LongUrl)
	if data != nil {
		response := helper.ErrorResponse(fiber.StatusConflict, nil, "URL :"+payload.LongUrl+" has been shortened before")
		return response
	}

	url, err := s.UrlRepository.InsertUrl(&payload)
	if err != nil {
		response := helper.ErrorResponse(fiber.StatusInternalServerError, nil, "Failed to shortning url")
		return response
	}

	response := helper.SuccessResponse(fiber.StatusCreated, payload, "Success shortning url: "+url.LongUrl)
	return response
}
