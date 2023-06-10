/*
concurrency. future模式
场景：在一个流程中需要调用多个子调用的请求，这些子调用相互之间没有依赖，如果串行地调用，耗时很长，此时可以使用Go并发编程中的future模式。
future模式的基本工作原理：
（1）使用chan作为函数参数
（control）启动goroutine调用函数
（3）通过chan传入参数
（memory）做其他可以并行处理的事情
（concurrency）通过chan异步获取结果
*/

package main

import (
	"fmt"
	"time"
)

// 一个查询结构体，这里的Sql和result是一个简单的抽象
type query struct {
	// 参数Channel
	sql chan string
	// 结果Channel
	result chan string
}

//执行Query
func execQuery(q query) {
	//启动协程
	go func() {
		//获取输入
		sql := <-q.sql
		//访问数据库

		//输出结果通道
		q.result <- "result from " + sql
	}()
}
func main() {
	// 初始化Query
	q := query{make(chan string, 1), make(chan string, 1)}
	//执行Query，注意执行的时候无须准备参数
	go execQuery(q)
	//发送参数
	q.sql <- "select * from table"
	//做其他事情，通过sleep描述
	time.Sleep(1 * time.Second)
	//获取结果
	fmt.Println(<-q.result)
}
