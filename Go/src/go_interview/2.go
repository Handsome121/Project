/*
写一段从在从channel中不断读取数据，并打印出来的程序。当channel被关闭时，
需要打印“Channel is closed”并结束程序。
*/
package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	listenCh := make(chan int)

	go readNum(ch, listenCh)

	for i := 0; i <= 4; i++ {
		ch <- i
		if i == 4 {
			close(ch)
		}
	}
	<-listenCh
}

func readNum(ch chan int, listenCh chan int) {
	for {
		if num, ok := <-ch; !ok {
			fmt.Println("channel closed!")
			listenCh <- 0
		} else {
			fmt.Println(num)
		}
	}
}
