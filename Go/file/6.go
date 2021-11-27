package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	file1Path := "E:/Project/Go/123.txt"
	file2Path := "E:/Project/Go/222.txt"
	content, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Printf("read file err=%v", err)
		return
	}
	_ = ioutil.WriteFile(file2Path, content, 0666)

}
