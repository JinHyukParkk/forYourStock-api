package routerv1

import (
	v1controller "forYourStock-api/controller/service/v1.1"

	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	{
		v1.GET("/ping", v1controller.Ping)
	}
}
