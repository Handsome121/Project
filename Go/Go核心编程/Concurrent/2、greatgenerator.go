package main

import (
	"fmt"
	"math/rand"
)

func GenerateIntA2() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}
func GenerateIntB2() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}
func GenerateInt2() chan int {
	ch := make(chan int, 20)
	go func() {
		for {
			//使用select的扇入技术（Fan in）增加生成的随机源
			select {
			case ch <- <-GenerateIntA2():
			case ch <- <-GenerateIntB2():
			}
		}
	}()
	return ch
}

func main() {
	ch := GenerateInt2()
	for i := 0; i < 100; i++ {
		fmt.Println(<-ch)
	}
}
