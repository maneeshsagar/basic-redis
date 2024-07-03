package service

import (
	"github.com/maneeshsagar/test/iolayer"
	"github.com/maneeshsagar/test/persistence"
)

type Driver interface {
	GetService(key string) (response *iolayer.GetResponse, code int, msg string, err error)
	PutService(request *iolayer.PulRequest) (code int, msg string, err error)
	PutTimoutService(request *iolayer.PutTimoutRequest) (code int, msg string, err error)
}

type RedisService struct {
	RedisPersistence persistence.Persistence
}

func NewRedisService() Driver {
	return &RedisService{
		RedisPersistence: persistence.NewRedisPersistence(),
	}
}
