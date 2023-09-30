package db

import (
	"database/sql"
	"time"

	"github.com/muhammadzhuhry/singkating-api-services/config"
)

func NewDB() *sql.DB {
	db, _ := sql.Open("mysql", config.MySQLConnectionURL)

	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
