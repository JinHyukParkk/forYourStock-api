package controller

import (
	"forYourStock/api"
	"os"
)

func GetList() {
	key := os.Getenv("APIKEY")
	api.GetCompanyList(key)
}