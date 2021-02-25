package router

import (
	"forYourStock-api/controller/service"
	routerv1 "forYourStock-api/router/v1.0"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

// 함수명 수정할 것
func SetRouter() {
	Router = gin.Default()

	test := Router.Group("/test")
	{
		test.GET("/v1", service.Test)
	}

	api := Router.Group("/api")
	{
		routerv1.SetRouter(api)
	}
}
