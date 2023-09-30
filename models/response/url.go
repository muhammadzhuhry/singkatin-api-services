package response

type Url struct {
	LongUrl      string `json:"longUrl"`
	ShortUrl     string `json:"shortUrl"`
	ShortenedUrl string `json:"shortenedUrl"`
	CreatedAt    string `json:"createdAt"`
	ExpiredAt    string `json:"expiredAt"`
}
