package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/maneeshsagar/test/service"
	"github.com/maneeshsagar/test/util"
)

func GetHandler(service service.Driver) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Query("key")
		response, code, msg, err := service.GetService(key)
		if err != nil {
			c.JSON(200, util.NewResponseWithModel(code, msg, nil))
			return
		}
		c.JSON(200, util.NewResponseWithModel(code, msg, response))
	}
}
