package routerv1

import (
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func SetRouter(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	{
		v1.GET("/ping", ping)
	}
}
