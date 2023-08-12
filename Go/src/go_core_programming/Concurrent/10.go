package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i // 向通道发送数据
		}
		close(ch) // 关闭通道->如果不关闭的话会造成死锁
	}()

	for num := range ch {
		fmt.Println(num)
	}

	fmt.Println("迭代完成")
}
