package db

import (
	"os"

	"github.com/go-redis/redis"
	"github.com/name5566/leaf/log"
)

// RedisC is opened redis client
var RSClient *redis.Client = nil

// InitRS init redis client
func InitRS() {
	if RSClient == nil {
		RSClient = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: os.Getenv("ROTA_REDISPASS"),
			DB:       0,
		})
	}
}

// CloseRS close redis client
func CloseRS() {
	if RSClient != nil {
		err := RSClient.Close()
		if err != nil {
			log.Fatal("%v", err)
		}
	}
}
