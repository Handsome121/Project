package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateIntA1() chan int {
	ch := make(chan int, 10)
	//启动一个goroutine用于生成随机数，函数返回一个通道用于获取随机数
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}
func main() {
	ch := GenerateIntA1()
	for {
		time.Sleep(5 * time.Second)
		fmt.Println(<-ch)
	}
}
