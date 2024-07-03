package persistence

import (
	"context"
	"fmt"
	"time"

	"github.com/maneeshsagar/test/db"
)

func (r *RedisPersistence) PutTimeoutPersistence(key string, value interface{}, timeout int) (err error) {
	client := db.GetRediClient()
	err = client.Set(context.TODO(), key, value, time.Second*time.Duration(timeout)).Err()
	if err != nil {
		fmt.Println("Error in PutTimeoutPersistence: ", err)
		return err
	}
	return err
}
