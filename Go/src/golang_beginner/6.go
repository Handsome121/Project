package main

import "fmt"

//func main() {
//	defer Fmt.Println("main end1")
//	defer Fmt.Println("main end2")
//	Fmt.Println("main:hello go 1")
//	Fmt.Println("main:hello go control")
//}

//func func1() {
//	Fmt.Println("A")
//}
//func func2() {
//	Fmt.Println("B")
//}
//func func3() {
//	Fmt.Println("C")
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
