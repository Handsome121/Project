package main

import "fmt"

//声明一种行的数据类型myint,是int的一个别名
type myint int
type Book struct {
	title string
	auth  string
}

func changeBook(book Book) {
	//传递一个book的副本,不会改变原来的book
	book.auth = "666"
}
func changeBook2(book *Book) {
	book.auth = "777"
}

func main() {
	var a myint = 10
	fmt.Println("a=", a)
	fmt.Printf("type of a = %T\n", a)
	var book Book
	book.title = "golang"
	book.auth = "zhangsan"
	fmt.Printf("%v\n", book)
	changeBook(book)
	fmt.Printf("%v\n", book)
	changeBook2(&book)
	fmt.Printf("%v\n", book)
}
