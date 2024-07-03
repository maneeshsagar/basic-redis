package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maneeshsagar/test/db"
	"github.com/maneeshsagar/test/routes"
)

// connect to the redis
// get values from the redis
// put values to the redis
// put values with timeout
// router to put values and get values : update the values
func main() {

	redis := db.InitializeAndGetRedisClient()
	defer redis.Close()
	r := gin.New()
	redisGroup := r.Group("/redis")
	routes.GetRoutes(redisGroup)
	r.Run(":8080")
}
