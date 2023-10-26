package main

import (
	"fmt"
	"sync"
)

func main() {
	//var wg sync1.WaitGroup
	//wg.Add(5)
	//out := []int{1, 2, 3, 4, 5}
	//for _, data := range out {
	//	data2 := data
	//	go func() {
	//		fmt.Printf("data is %v\n", data2)
	//		wg.Done()
	//	}()
	//}
	//wg.Wait()

	var wg sync.WaitGroup
	wg.Add(5)
	out := []int{1, 2, 3, 4, 5}
	for _, data := range out {
		go func(data int) {
			fmt.Printf("data is %v\n", data)
			wg.Done()
		}(data)
	}
	wg.Wait()
}
