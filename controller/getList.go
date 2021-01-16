package controller

import (
	"forYourStock/api"
	"os"
)

func GetList() {
	key := os.Getenv("APIKEY")

	// 회사 리스트 list.zip 파일 불러오기
	api.GetCompanyList(key)

	// 압출풀기
	api.GetList()
}