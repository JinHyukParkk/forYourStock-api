package router

import (
	routertest "forYourStock-api/router/test"
	routerv1 "forYourStock-api/router/v1.0"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

// 함수명 수정할 것
func SetRouter() {
	Router = gin.Default()

	test := Router.Group("/test")
	{
		routertest.SetRouter(test)
	}

	api := Router.Group("/api")
	{
		routerv1.SetRouter(api)
	}
}
