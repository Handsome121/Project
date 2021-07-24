package main

import "fmt"

type Reader interface {
	ReaderBook()
}
type Writer interface {
	WriterBook()
}

//具体类型
type Book struct {
}

func (this *Book) ReaderBook() {
	fmt.Println("Read a book")
}
func (this *Book) WriterBook() {
	fmt.Println("Write a book")
}
func main() {
	b := &Book{} //b:pair<type:Book,value:book{}地址>
	var r Reader //r:pair<type:,value:>
	r = b        //r:pair<type:Book,value:book{}地址>
	r.ReaderBook()
	var w Writer
	w = r.(Writer)
	w.WriterBook()
}
