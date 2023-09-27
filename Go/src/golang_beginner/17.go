package main

import "fmt"

//interface{}是万能数据类型
func myFunc(args interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(args)
	//interface{}该如何区分此时引用的数据类型到底是什么？
	//给interface{}提供“类型断言”机制
	value, ok := args.(string)
	if !ok {
		fmt.Println("args is not string type")
	} else {
		fmt.Println("args is string type,value=", value)
		fmt.Printf("value type is %T\n", value)
	}
}

type Book struct {
	auth string
}

func main() {
	book := Book{"zhang3"}
	myFunc(book)
	myFunc(200)
	myFunc("123")
}
