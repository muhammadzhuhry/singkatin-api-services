package entity

import "time"

type Url struct {
	ID        int       `json:"id"`
	LongUrl   string    `json:"longUrl"`
	ShortUrl  string    `json:"shortUrl"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiredAt time.Time `json:"expiredAt"`
}
