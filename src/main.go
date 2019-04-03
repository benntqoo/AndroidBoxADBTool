package main

import (
	"adbHelper/src/activity"
	"fmt"
	"os"
)

const signedFileName string = "signed.cfg"
const projectFileName string = "project.cfg"

func main() {
	//讀取 singed 檔案資料
	signedStr, err := os.Open(signedFileName)
	if err != nil {
		fmt.Println("signed file err ", err)
	}
	fmt.Println(signedStr)

	//讀取 project 檔案資料
	projectStr, err := os.Open(projectFileName)
	if err != nil {
		fmt.Println("project file err ", err)
	}
	fmt.Println(projectStr)

	activity.CreateApp()
}
