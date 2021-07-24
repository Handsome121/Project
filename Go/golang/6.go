package main

import "fmt"

//func main() {
//	defer fmt.Println("main end1")
//	defer fmt.Println("main end2")
//	fmt.Println("main:hello go 1")
//	fmt.Println("main:hello go 2")
//}

//func func1() {
//	fmt.Println("A")
//}
//func func2() {
//	fmt.Println("B")
//}
//func func3() {
//	fmt.Println("C")
//}
//func main() {
//	defer func1()
//	defer func2()
//	defer func3()
//}

func deferFunction() int {
	fmt.Println("defer function called...")
	return 0
}
func returnFunction() int {
	fmt.Println("return function called...")
	return 0
}
func returnAndDefer() int {
	defer deferFunction()
	return returnFunction()
}
func main() {
	returnAndDefer()
}
