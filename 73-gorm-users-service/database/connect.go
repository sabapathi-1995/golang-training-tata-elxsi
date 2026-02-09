package database

import (
	"log/slog"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Retries        int = 10
	RETRY_DURATION     = time.Second * 5
)

func ConnectDb(dsn string) (*gorm.DB, error) {

	count := 1
retry:
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		if count == Retries {
			return db, err
		}
		slog.Error(err.Error(), "--> retrying it for the ", count)
		time.Sleep(RETRY_DURATION)
		count++
		goto retry
	}
	return db, err
}

//dsn := "host=localhost user=postgres password=postgres dbname=usersdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
