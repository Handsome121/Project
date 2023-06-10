//有时希望生成器能够自动退出，可以借助Go通道的退出通知机制（close channel to broadcast）实现。
package main

import (
"fmt"
"math/rand"
)

func GenerateIntA3(done chan struct{}) chan int {
	ch := make(chan int)
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

func main() {
	done := make(chan struct{})
	ch := GenerateIntA3(done)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//不再需要生成器，通过close chan发送一个通知给生成器
	close(done)
	for v := range ch {
		fmt.Println(v)
	}
}
