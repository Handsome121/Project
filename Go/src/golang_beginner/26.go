package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	/*for {
		if data, ok := <-c; ok { //关闭之后仍然能从channel中获取数据，等到没有数据可以取了，ok才会返回false
			Fmt.Println(data)
		} else {
			break
		}
	}*/
	for data := range c {
		fmt.Println(data)
	}
	fmt.Println("main finish")
}
