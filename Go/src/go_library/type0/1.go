package main

import "fmt"

type MyFunc func(int, int) int

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func main() {
	var f MyFunc
	f = add
	result := f(3, 4)
	fmt.Println(result)
	f = subtract
	result = f(7, 2)
	fmt.Println("Subtraction result:", result)
}
