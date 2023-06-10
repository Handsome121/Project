package main

import "fmt"

func main() {
	var a [][]int
	for i := 0; i < 10; i++ {
		var tmp []int
		for j := 0; j < 10; j++ {
			tmp = append(tmp, j)
		}
		a = append(a, tmp)
	}
	fmt.Println(a)
}
