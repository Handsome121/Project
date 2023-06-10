package main

import "fmt"

func main() {
	var out []int
	for i := 0; i < 3; i++ {
		//j := i
		//fmt.Println(&i)
		out = append(out, i)
	}
	fmt.Println("values:", out[0], out[1], out[2])

	//注意:如果需要以指针的形式存放循环变量，则可以显式地拷贝一次
}
