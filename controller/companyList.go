package controller

import (
	"fmt"
	"forYourStock/controller/service"
	"os"
)

func GetList() {
	key := os.Getenv("APIKEY")

	// 회사 리스트 list.zip 파일 불러오기
	service.GetCompanyList(key)

	// 압출풀기
	go service.GetList()
	fmt.Println("Call GetList")
}
