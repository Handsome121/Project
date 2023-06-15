package main

import "fmt"

func SelectForChan(c chan string) {
	var recv string
	send := "hello"

	select {
	case recv = <-c:
		fmt.Printf("received %s \n", recv)
	case c <- send:
		fmt.Printf("sent %s\n", send)
	}
}

//select 拥有两个case语句，分别对应管道的读操作和写操作，至于最终执行哪个case语句，取决于函数传入的管道

func main() {
	//第一种情况，管道没有缓冲区，此时管道既不能读，也不能写，两个case语句均不执行，select陷入阻塞
	//c := make(chan string)
	//SelectForChan(c)

	//第二种情况，管道又缓冲区且还可以存放至少一个数据,此时管道可以写入，写操作的case可以得到执行，且结束后函数退出
	//c := make(chan string, 1)
	//SelectForChan(c)

	//管道有缓冲区且缓冲区已放满数据,此时管道可以读取，读操作对应的case语句得到执行，且执行结束后函数退出
	//c := make(chan string, 1)
	//c <- "hello"
	//SelectForChan(c)

	//管道有缓冲区，缓冲区中已有部分数据且还可以存入数据,select将随机挑选一个case语句执行，任意一个case语句执行结束后函数就退出
	c := make(chan string, 2)
	c <- "hello"
	SelectForChan(c)

}
