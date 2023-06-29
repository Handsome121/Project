///*
//3. 每个请求一个goroutine
//这种并发模式就是来一个请求或任务就启动一个goroutine去处理，典型的是Go中的HTTP server服务。源码如下：
//*/
//
package main

import (
	"fmt"
	"sync"
)

// 工作任务
type task1 struct {
	begin  int
	end    int
	result chan<- int
}
//
//任务执行：计算begin到end的和，执行结果写入结果chan result
func (t *task1) do() {
	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum += i
	}
	t.result <- sum
}
func main() {
	// 创建任务通道
	taskChan := make(chan task, 10)
	// 创建结果通道
	resultChan := make(chan int, 10)
	// wait用于同步等待任务的执行
	wait := &sync.WaitGroup{}

	//初始化task的goroutine，计算100个自然数之和
	go InitTask1(taskChan, resultChan, 100)
	// 每个task启动一个goroutine进行处理
	go DistrubuteTask1(taskChan, wait, resultChan)
	//通过结果通道获取结果并汇总
	sum := ProcessResult1(resultChan)
	fmt.Println("sum=", sum)
}

// 构建task并写入task通道
func InitTask1(taskChan chan<- task, r chan int, p int) {
	qu := p / 10
	mod := p % 10
	high := qu * 10
	for j := 0; j < qu; j++ {
		b := 10*j + 1
		e := 10 * (j + 1)
		tsk := task{
			begin:  b,
			end:    e,
			result: r,
		}
		taskChan <- tsk
	}
	if mod != 0 {
		tsk := task{
			begin:  high + 1,
			end:    p,
			result: r,
		}
		taskChan <- tsk
	}
	close(taskChan)
}

//读取task chan,每个task启动一个worker goroutine进行处理，并等待每个task运行完，关闭结果通道
func DistrubuteTask1(taskChan chan task, wait *sync.WaitGroup, result chan int) {
	for v := range taskChan {
		wait.Add(1)
		go ProcessTask(v, wait)
	}
	wait.Wait()
	close(result)
}

//goroutine处理具体工作，并将结果发送到结果通道
func ProcessTask1(t task, wait *sync.WaitGroup) {
	t.do()
	wait.Done()
}

//读取结果通道，汇总结果
func ProcessResult1(resultChan chan int) int {
	sum := 0
	for r := range resultChan {
		sum += r
	}
	return sum
}
