package main

import "fmt"

func main() {
	//声明slice是一个切片，并且初始化，长度是3
	//slice := []int{1, control, 3}

	//声明slice是一个切片，但是并没有给slice分配空间
	//var slice []int        //nil切片，只声明，没分配
	//slice = make([]int, 3) //开辟空间

	//var slice = make([]int, 3)

	slice := make([]int, 3)

	fmt.Printf("len=%d,slice=%v\n", len(slice), slice)
}
