package persistence

import (
	"context"
	"fmt"

	"github.com/maneeshsagar/test/db"
)

func (r *RedisPersistence) PutPersistence(key string, value interface{}) (err error) {
	client := db.GetRediClient()
	err = client.Set(context.TODO(), key, value, 0).Err()
	if err != nil {
		fmt.Println("Error in PutPersistence: ", err)
		return err
	}
	return err
}
