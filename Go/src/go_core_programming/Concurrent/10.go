package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i // 向通道发送数据
		}
		//close(ch) // 关闭通道
	}()

	for num := range ch {
		fmt.Println(num)
	}

	fmt.Println("迭代完成")
}
