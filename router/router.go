package router

import (
	"fmt"
	"forYourStock-api/controller"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func SetRouter() {
	Router = gin.Default()

	test := Router.Group("/test")
	{
		test.GET("/v1", func(c *gin.Context) {
			controller.GetList()

			c.JSON(200, gin.H{
				"msg": "good",
			})

			fmt.Println("Return test/v1")
		})
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
