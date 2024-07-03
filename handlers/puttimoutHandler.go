package handlers

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/maneeshsagar/test/iolayer"
	"github.com/maneeshsagar/test/service"
	"github.com/maneeshsagar/test/util"
)

func PutTimoutHandler(service service.Driver) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody, err := io.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println("Error reading the request body")
			c.JSON(500, util.NewResponseWithModel(401, "Error reading the request body", nil))
			return
		}

		request := new(iolayer.PutTimoutRequest)

		err = json.Unmarshal(requestBody, request)
		if err != nil {
			fmt.Println("request body is not as expected")
			c.JSON(500, util.NewResponseWithModel(401, "request body is not as expected", nil))
			return
		}

		code, msg, err := service.PutTimoutService(request)
		if err != nil {
			fmt.Println("Error in PutHandler: ", err)
			c.JSON(200, util.NewResponseWithModel(code, msg, nil))
			return
		}

		c.JSON(200, util.NewResponseWithModel(code, msg, nil))
	}
}
