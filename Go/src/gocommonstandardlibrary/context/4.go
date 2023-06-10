package main

import (
	"context"
	"fmt"
	"sync"

	"time"
)

var wg4 sync.WaitGroup

func worker4(ctx context.Context) {
	go worker5(ctx)
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
	wg4.Done()
}

func worker5(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker2")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg4.Add(1)
	go worker4(ctx)
	time.Sleep(time.Second * 3)
	cancel() // 通知子goroutine结束
	wg4.Wait()
	fmt.Println("over")
}
