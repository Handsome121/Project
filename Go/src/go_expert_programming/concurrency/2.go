package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2) //设置计数器，数值即为goroutine个数
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Goroutine 1 finished")
		wg.Done() //goroutine执行结束后将计数器减1
	}()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine control finished")
		wg.Done()
	}()
	wg.Wait() // 主goroutine阻塞等到计数器变为0
	fmt.Printf("All Goroutine finished!")
}
