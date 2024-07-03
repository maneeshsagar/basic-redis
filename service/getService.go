package service

import (
	"fmt"

	"github.com/maneeshsagar/test/iolayer"
	"github.com/redis/go-redis/v9"
)

func (r *RedisService) GetService(key string) (response *iolayer.GetResponse, code int, msg string, err error) {

	value, err := r.RedisPersistence.GetPersistence(key)
	if err != nil && err != redis.Nil {
		fmt.Println("Error in GetService: ", err)
		return response, 500, "Internal Server Error", err
	}

	if err == redis.Nil {
		return nil, 404, "record Not Found", err
	}

	response = new(iolayer.GetResponse)
	response.Key = key
	response.Value = value
	return response, 200, "Success", nil
}
