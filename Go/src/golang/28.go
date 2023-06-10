package main

import "fmt"

func myFunc(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("data = ", data)
	}

}

func test(args ...int) {
	myFunc(args...)

}

func main() {
	//test(1, 2, 3, 4)
	test([]int{1, 2, 3, 4}...)
}
