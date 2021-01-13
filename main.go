package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func setRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello gin")
	})

	return r
}

func main() {
	r := setRouter()

	r.Run(":8080")
}
