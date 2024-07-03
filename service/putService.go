package service

import (
	"fmt"

	"github.com/maneeshsagar/test/iolayer"
)

func (r *RedisService) PutService(request *iolayer.PulRequest) (code int, msg string, err error) {

	err = r.RedisPersistence.PutPersistence(request.Key, request.Value)
	if err != nil {
		fmt.Println("Error in PutService: ", err)
		return 500, "Internal Server Error", err
	}

	return 200, "Success", nil
}
