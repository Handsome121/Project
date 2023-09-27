package main

import "fmt"

func main() {
	var a string
	//pair<type:string,value:"aceld">
	a = "aceld"
	fmt.Println(a)
	var allType interface{}
	allType = a
	fmt.Println(allType)
}
