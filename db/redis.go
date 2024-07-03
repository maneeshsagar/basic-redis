package db

import (
	"fmt"
	"sync"

	"github.com/maneeshsagar/test/config"
	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	redisOnce   sync.Once
)

func InitializeAndGetRedisClient() *redis.Client {
	redisOnce.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     config.RedisSettings.Addr,
			Password: config.RedisSettings.Password,
			DB:       config.RedisSettings.DB,
		})
	})
	fmt.Println("Redis client initialized")
	return redisClient
}

func GetRediClient() *redis.Client {
	if redisClient != nil {
		return redisClient
	}

	redisClient = InitializeAndGetRedisClient()
	return redisClient
}
