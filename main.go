package main

import (
	"forYourStock-api/router"
)

func main() {
	router.SetRouter()

	router.Router.Run(":8080")
}
