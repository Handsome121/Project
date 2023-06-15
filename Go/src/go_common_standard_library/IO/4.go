package main

import (
	"fmt"
	"io/ioutil"
)

/*ioutil工具包*/

func wr1() {
	err := ioutil.WriteFile("./yyy.txt", []byte("www.5lmh.com"), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func re1() {
	content, err := ioutil.ReadFile("./yyy.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}

func main() {
	wr1()
	//re1()
}
