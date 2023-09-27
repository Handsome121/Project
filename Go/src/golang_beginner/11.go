package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}
	s1 := s[0:2]
	fmt.Println(s1)
	s1[0] = 6
	fmt.Println(s)
	fmt.Println(s1) //切片赋值与Python中不一样
	//copy可以将底层数组的slice一起进行拷贝
	s2 := make([]int, 4)
	copy(s2, s)
	s2[0]=0
	fmt.Println(s)
	fmt.Println(s2)


}
