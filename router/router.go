package router

import (
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func SetRouter() {
	Router = gin.Default()

	api := Router.Group("/test")
	{
		api.GET("/v1", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "good",
			})
		})
	}


}