package main

import (
	"fmt"
	"sync"

	"time"
)

var wg1 sync.WaitGroup

// 初始的例子

func worker1() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
	}
	// 如何接收外部命令实现退出
	wg1.Done()
}

func main() {
	wg.Add(1)
	go worker1()
	// 如何优雅的实现结束子goroutine
	wg1.Wait()
	fmt.Println("over")
}
