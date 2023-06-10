package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//官方版答案
var wg3 sync.WaitGroup

func worker3(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
	wg3.Done()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg3.Add(1)
	go worker3(ctx)
	time.Sleep(time.Second * 3)
	cancel() //通知子Goroutine结束
	wg3.Wait()
	fmt.Println("over")
}
