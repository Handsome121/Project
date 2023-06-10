package main

import "fmt"

func printArray(myArray []int) {
	for _, value := range myArray {
		fmt.Println("value=", value)
	}
}
func main() {
	//动态数组
	myArray := []int{1, 2, 3, 4}
	fmt.Printf("myArray type is %T\n", myArray)
	printArray(myArray)
}
