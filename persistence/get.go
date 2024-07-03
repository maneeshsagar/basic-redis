package persistence

import (
	"context"
	"fmt"

	"github.com/maneeshsagar/test/db"
)

func (r *RedisPersistence) GetPersistence(key string) (value string, err error) {
	client := db.GetRediClient()
	value, err = client.Get(context.TODO(), key).Result()
	if err != nil {
		fmt.Println("Error in GetPersistence: ", err)
		return value, err
	}

	return value, err
}
