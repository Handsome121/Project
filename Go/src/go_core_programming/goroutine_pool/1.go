package main

import (
	"fmt"
	"sync"
)

// Task 任务结构体
type Task struct {
	ID  int
	Job func()
}

// Pool 协程池结构体
type Pool struct {
	taskQueue chan Task
	wg        sync.WaitGroup
}

// NewPool 创建协程池
func NewPool(numWorkers int) *Pool {
	p := &Pool{
		taskQueue: make(chan Task),
	}

	p.wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go p.worker()
	}

	return p
}

// AddTask 添加任务到协程池
func (p *Pool) AddTask(task Task) {
	p.taskQueue <- task
}

// 工作协程
func (p *Pool) worker() {
	for task := range p.taskQueue {
		fmt.Printf("Worker %d started task %d\n", task.ID, task.ID)
		task.Job()
		fmt.Printf("Worker %d finished task %d\n", task.ID, task.ID)
	}
	p.wg.Done()
}

// Wait 等待所有任务完成
func (p *Pool) Wait() {
	close(p.taskQueue)
	p.wg.Wait()
}

func main() {
	// 创建一个协程池，设置工作协程数为3
	pool := NewPool(3)

	// 添加任务到协程池
	for i := 0; i < 10; i++ {
		taskID := i
		task := Task{
			ID: taskID,
			Job: func() {
				fmt.Printf("Task %d is running\n", taskID)
			},
		}
		pool.AddTask(task)
	}

	// 等待所有任务完成
	pool.Wait()
}
