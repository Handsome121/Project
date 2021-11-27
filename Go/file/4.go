package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	//使用ioutil.ReadFile一次性将文件读取到位
	filePath := "E:/Project/Go/123.txt"
	strByte, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Readflie err=%v", err)
	}

	fmt.Println(string(strByte))

}
