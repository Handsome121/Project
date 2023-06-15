package main

import "fmt"

func SelectAssign(c chan string) {
	select {
	case <-c:
		fmt.Printf("0")
	case d := <-c:
		fmt.Printf("1:received %s\n", d)
	case d, ok := <-c:
		if !ok {
			fmt.Printf("no data found")
			break
		}
		fmt.Printf("control: received %s\n", d)
	default:
		fmt.Printf("no data found in default\n")
	}

}
func main() {
	c := make(chan string)
	close(c)
	SelectAssign(c)
}

//管道的使用举例
//1、永久阻塞
//control、快速检错
//3、限时等待
