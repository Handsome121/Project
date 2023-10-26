/*
golang两个协程，要求实现交替打印的效果，go1 => 1 3 5 7；go2 => 2 4 6 8 ；整体输出是 go1和go2交替打印。
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	wg:=sync.WaitGroup{}
	wg.Add(2)

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 1; i <= 7; i += 2 {
			ch1 <- i
		}
		close(ch1)
		wg.Done()
	}()

	go func() {
		for i := 2; i <= 8; i += 2 {
			ch2 <- i
		}
		close(ch2)
		wg.Done()
	}()
	
	for {
		select {
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				break
			}
			fmt.Println("go1 =>", v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				break
			}
			fmt.Println("go2 =>", v)
		}
		if ch1 == nil && ch2 == nil {
			break
		}
	}
}

