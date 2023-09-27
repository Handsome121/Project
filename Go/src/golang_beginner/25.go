package main

import (
	"fmt"
	"time"
)

func main() {
	/*go func() {
		defer Fmt.Println("A.defer")
		func() {
			defer Fmt.Println("B.defer")
			runtime.Goexit()
			Fmt.Println("B")
		}()
		Fmt.Println("A")
	}()*/
	go func(a int, b int) bool {
		fmt.Println("a=", a, "b=", b)
		return true
	}(10, 20)
	//死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
