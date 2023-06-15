package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	file := "E:/Project/Go/123.txt"
	content, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Printf("read gofile error=%v\n", err)
	}
	fmt.Println(string(content))
}
