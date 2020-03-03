package db

import (
	"rota/conf"
	"time"

	"github.com/go-redis/redis"
	"github.com/name5566/leaf/log"
)

// RedisC is opened redis client
var RSClient *redis.Client = nil

// InitRS init redis client
func InitRS() {
	if RSClient == nil {
		for {
			RSClient = redis.NewClient(&redis.Options{
				Addr:     conf.Server.RedisAddr,
				Password: conf.Server.RedisPass,
				DB:       0,
			})
			_, err := RSClient.Ping().Result()
			if err == nil {
				log.Release("Redis connected")
				break
			} else {
				log.Error("%v", err)
				time.Sleep(5 * time.Second)
			}
		}
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
