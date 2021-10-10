package main

import "fmt"

func AppendDemo() {
	x := make([]int, 0, 10)
	x = append(x, 1, 2, 3)
	y := append(x, 4)
	z := append(x, 5)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
func main() {
	AppendDemo()
}

//一般情况下，使用append向切片追加新的元素时，都会用原切片变量接受返回值来获得更新
//如果使用心的变量接受返回值，则需要考虑append返回的切片是否跟原切片共享底层的数组
