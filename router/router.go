package router

import (
	"github.com/gin-gonic/gin"
	"forYourStock-api/controller"
	"fmt"
)

var Router *gin.Engine

func SetRouter() {
	Router = gin.Default()

	api := Router.Group("/test")
	{
		api.GET("/v1", func(c *gin.Context) {
			controller.GetList()
			
			c.JSON(200, gin.H{
				"msg": "good",
			})

			fmt.Println("Return test/v1")
		})
	}


}