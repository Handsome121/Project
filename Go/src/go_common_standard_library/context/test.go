package main

import "fmt"

type MyInt int

func main() {
	fmt.Println(MyInt(123))
	fmt.Printf("类型是%T", MyInt(123))
}
