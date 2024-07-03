package service

import "github.com/maneeshsagar/test/iolayer"

func (r *RedisService) PutTimoutService(request *iolayer.PutTimoutRequest) (code int, msg string, err error) {
	err = r.RedisPersistence.PutTimeoutPersistence(request.Key, request.Value, request.Timeout)
	if err != nil {
		return 500, "Internal Server Error", err
	}
	return 200, "Success", nil
}
