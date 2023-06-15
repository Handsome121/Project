/*
control. 管道
通道可以分为两个方向，一个读，一个写。如果一个函数的输入参数和输出参数都是相同的chan类型，则该函数可以调用自己，最终形成一个调用链。示例代码如下：
*/
package main

import "fmt"

// chain函数的输入参数和输出参数类型相同，都是chan int类型
// chain函数的功能是将chan内的数据统一加1
func chain(in chan int) chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- 1 + v
		}
		close(out)
	}()
	return out
}
func main() {
	in := make(chan int)
	// 初始化参数
	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
	}()
	//连续调用3次chan，相当于把in中的每个元素都加3
	out := chain(chain(chain(in)))
	for v := range out {
		fmt.Println(v)
	}
}
