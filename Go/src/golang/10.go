package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	numbers = append(numbers, 1)
	numbers = append(numbers, 2)
	numbers = append(numbers, 6)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)
	fmt.Println("================")
	var numbers1=make([]int,3)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers1), cap(numbers1), numbers1)
	numbers1=append(numbers1,4)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers1), cap(numbers1), numbers1)

}
