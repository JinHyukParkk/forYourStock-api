package main

import (
	"github.com/gin-gonic/gin"
	"forYourStock-api/router"
	"os"
	"io"
)

func main() {
	 // Disable Console Color, you don't need console color when writing the logs to file.
	 gin.DisableConsoleColor()

	 // Logging to a file.
	 f, _ := os.Create("log/gin.log")
	 gin.DefaultWriter = io.MultiWriter(f)

	router.SetRouter()

	router.Router.Run(":8080")
}
