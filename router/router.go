package router

import (
	"forYourStock-api/controller/service"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func SetRouter() {
	Router = gin.Default()

	test := Router.Group("/test")
	{
		test.GET("/v1", service.Test)
	}

	api := Router.Group("/")
	{
		api.GET("/v1", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "api good",
			})
		})
	}
}
