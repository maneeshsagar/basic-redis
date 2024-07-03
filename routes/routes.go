package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/maneeshsagar/test/handlers"
	"github.com/maneeshsagar/test/service"
)

func GetRoutes(r *gin.RouterGroup) {

	redisService := service.NewRedisService()

	r.PUT("/put", handlers.PutHandler(redisService))
	r.GET("/get", handlers.GetHandler(redisService))
	r.PUT("/puttimeout", handlers.PutTimoutHandler(redisService))

}
