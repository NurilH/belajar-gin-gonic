package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Address  string
	Password string
	DB       int
}

func NewRedisConnection(c *Redis) (rdb *redis.Client, err error) {
	// Membuat context
	ctx := context.Background()

	rdb = redis.NewClient(&redis.Options{
		Addr:     c.Address,
		Password: c.Password,
		DB:       c.DB,
	})

	// Tes koneksi
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Invalid connect to Redis: %v", err)
		return nil, err
	}
	return
}
