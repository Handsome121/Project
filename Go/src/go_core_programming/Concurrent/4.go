/*
一个融合了并发、缓冲、退出通知等多重特性的生成器。
*/
package main

import (
	"fmt"
	"math/rand"
)

// GenerateIntA4 done接收通知退出信号
func GenerateIntA4(done chan struct{}) chan int {
	ch := make(chan int, 5)
	go func() {
	Label:
		for {
			// 通过select监听一个信号chan来确定是否停止生成
			select {
			case ch <- rand.Int():
			case <-done:
				break Label
			}
		}
		close(ch)
	}()
	return ch
}

// GenerateIntB4 done接收通知退出信号
func GenerateIntB4(done chan struct{}) chan int {
	ch := make(chan int, 10)
	go func() {
	Label:
		for {
			// 通过select监听一个信号chan来确定是否停止生成
			select {
			case ch <- rand.Int():
			case <-done:
				break Label
			}
		}
		close(ch)
	}()
	return ch
}

// GenerateInt 通过select执行扇入（Fan in )操作
func GenerateInt(done chan struct{}) chan int {
	ch := make(chan int)
	send := make(chan struct{})
	go func() {
	Label:
		for {
			//使用select的扇入技术（Fan in）增加生成的随机源
			select {
			case ch <- <-GenerateIntA4(send):
			case ch <- <-GenerateIntB4(send):
			case <-done:
				send <- struct{}{}
				send <- struct{}{}
				break Label

			}
		}
		close(ch)
	}()
	return ch
}

func main() {
	// 构建一个作为接收退出信号的chan
	done := make(chan struct{})
	//启动生成器
	ch := GenerateInt(done)
	//获取生成器资源
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	//通知生产者停止生成
	done <- struct{}{}
	fmt.Println("stop generate")
}
