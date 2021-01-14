package main

import (
	"forYourStock-api-server/router"
)

func main() {
	router.SetRouter()

	router.Router.Run(":8080")
}
