package main

import "fmt"

func main() {
	//1、使用append()函数时，谨记append可能会产生新的切片，并谨慎地处理返回值
	//2、使用append()函数时，谨记append可能会追加nil值，应该尽量避免追加无意义的元素

	x := make([]int, 0, 10)
	x = append(x, 1, 2, 3)
	y := append(x, 4)
	z := append(x, 5)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
