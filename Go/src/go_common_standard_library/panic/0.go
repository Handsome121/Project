package main

import (
	"fmt"
	"time"
)

//1、defer 需要放在 panic 之前定义，另外recover只有在 defer 调用的函数中才有效.
//2、recover处理异常后，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后的那个点.
//3、多个 defer 会形成 defer 栈，后定义的 defer 语句会被最先调用.
//panic (主动爆出异常) 与 recover （收集异常）.
//recover 用来对panic的异常进行捕获. panic 用于向上传递异常，执行顺序是在 defer 之后.

func main() {
	f()
	fmt.Println("end")
}

func f() {
	defer func() { //必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("2")
		if err := recover(); err != nil {
			fmt.Println(err) //这里的err其实就是panic传入的内容，bug
		}
		fmt.Println("3")
	}()
	for {
		fmt.Println("1")
		a := []string{"a", "b"}
		fmt.Println(a[3]) // 这里slice越界异常了
		//panic("bug")
		fmt.Println("4")
		time.Sleep(1 * time.Second)
	}
}
