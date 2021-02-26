package routertest

import (
	testcontroller "forYourStock-api/controller/service/test"

	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.RouterGroup) {
	v1 := router.Group("/")
	{
		v1.GET("/", testcontroller.Test)
	}
}
