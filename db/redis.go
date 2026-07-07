package db

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(ctx context.Context) *redis.Client {
	dbAddress := os.Getenv("REDIS_ADDRESS")
	dbPassword := os.Getenv("REDIS_PASSWORD")
	dbUsername := os.Getenv("REDIS_USERNAME")
	if dbAddress == "" || dbPassword == "" || dbUsername == "" {
		panic("Redis configuration is missing. Please set REDIS_ADDRESS, REDIS_PASSWORD, and REDIS_USERNAME environment variables.")
	}

	db := redis.NewClient(&redis.Options{
		Addr:          dbAddress,
		Username:      dbUsername,
		Password:      dbPassword,
		ClientName:    "my-redis-client",
		DB:            0,
		DialTimeout:   5 * time.Second,
		DialerRetries: 3,
	})

	return db
}
