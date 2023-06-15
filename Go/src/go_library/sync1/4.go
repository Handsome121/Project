/*
sync1.Once：用于只执行一次的操作。
*/
package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	fmt.Println("Initializing...")
}

func main() {
	for i := 0; i < 5; i++ {
		once.Do(initialize)
		fmt.Println("Do something else...")
	}
}
