package repository

import (
	"database/sql"

	"github.com/muhammadzhuhry/singkating-api-services/models/entity"
)

type UrlRepository struct {
	DB *sql.DB
}

func NewUrlRepository(db *sql.DB) *UrlRepository {
	return &UrlRepository{
		DB: db,
	}
}

func (r *UrlRepository) InsertUrl(url *entity.Url) (*entity.Url, error) {
	command := "INSERT INTO urls(long_url, short_url, expired_at) VALUES (?, ?, ?)"
	result, err := r.DB.Exec(command, url.LongUrl, url.ShortUrl, url.ExpiredAt)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	url.ID = int(id)
	return url, nil
}