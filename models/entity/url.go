package entity

type Url struct {
	ID        int    `json:"id"`
	LongUrl   string `json:"longUrl"`
	ShortUrl  string `json:"shortUrl"`
	CreatedAt string `json:"createdAt"`
	ExpiredAt string `json:"expiredAt"`
}
