package service

import (
	"fmt"
	"forYourStock-api/controller"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	controller.GetList()

	c.JSON(200, gin.H{
		"msg": "good",
	})

	fmt.Println("Return test/v1")
}
