package service

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/muhammadzhuhry/singkatin-api-services/config"
	"github.com/muhammadzhuhry/singkatin-api-services/helper"
	"github.com/muhammadzhuhry/singkatin-api-services/models"
	"github.com/muhammadzhuhry/singkatin-api-services/models/entity"
	"github.com/muhammadzhuhry/singkatin-api-services/models/request"
	"github.com/muhammadzhuhry/singkatin-api-services/models/response"
	"github.com/muhammadzhuhry/singkatin-api-services/repository"
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
	data, _ := s.UrlRepository.FindExistedUrl(payload.LongUrl, "long")
	if data != nil {
		response := helper.ErrorResponse(fiber.StatusConflict, nil, "URL :"+payload.LongUrl+" has been shortened before")
		return response
	}

	url, err := s.UrlRepository.InsertUrl(&payload)
	if err != nil {
		response := helper.ErrorResponse(fiber.StatusInternalServerError, nil, "Failed to shortning url")
		return response
	}

	resultUrl := response.Url{
		LongUrl:      url.LongUrl,
		ShortUrl:     url.ShortUrl,
		ShortenedUrl: config.Domain + url.ShortUrl,
		CreatedAt:    url.CreatedAt,
		ExpiredAt:    url.ExpiredAt,
	}
	response := helper.SuccessResponse(fiber.StatusCreated, resultUrl, "Success shortning url: "+resultUrl.LongUrl)
	return response
}

func (s *UrlService) RedirectUrl(c *fiber.Ctx, request string) models.Response {
	location := time.FixedZone(config.Timezone, 7*60*60)
	url, err := s.UrlRepository.FindExistedUrl(request, "short")
	if err != nil {
		response := helper.ErrorResponse(fiber.StatusNotFound, nil, "Cannot find url : "+request)
		return response
	}

	expired, _ := time.ParseInLocation("2006-01-02 15:04:05", url.ExpiredAt, location)
	if expired.Before(time.Now()) {
		s.UrlRepository.DeleteExpiredUrl(url.ShortUrl)
		response := helper.ErrorResponse(fiber.StatusBadRequest, nil, "url: "+request+" is expired")
		return response
	}

	result := response.Url{
		LongUrl:      url.LongUrl,
		ShortUrl:     url.ShortUrl,
		ShortenedUrl: config.Domain + url.ShortUrl,
		CreatedAt:    url.CreatedAt,
		ExpiredAt:    url.ExpiredAt,
	}
	response := helper.SuccessResponse(fiber.StatusCreated, result, "Success get long url: "+result.ShortenedUrl)
	return response
}
