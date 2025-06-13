package task_scheduler

import (
	"context"
	"errors"
	"sync"
	"time"
)

// Task 表示一个可执行的任务
type Task interface {
	Execute(ctx context.Context) error
	Name() string
}

// Scheduler 任务调度器
type Scheduler struct {
	taskQueue    chan Task
	workerPool   chan struct{}
	wg           sync.WaitGroup
	shutdownChan chan struct{}
	timeout      time.Duration
}

// NewScheduler 创建新的调度器
func NewScheduler(workerCount int, queueSize int, timeout time.Duration) *Scheduler {
	return &Scheduler{
		taskQueue:    make(chan Task, queueSize),
		workerPool:   make(chan struct{}, workerCount),
		shutdownChan: make(chan struct{}),
		timeout:      timeout,
	}
}

// Start 启动调度器
func (s *Scheduler) Start() {
	for {
		select {
		case <-s.shutdownChan:
			return
		case task := <-s.taskQueue:
			s.wg.Add(1)
			s.workerPool <- struct{}{} // 获取worker槽位
			go s.executeTask(task)
		}
	}
}

// executeTask 执行单个任务
func (s *Scheduler) executeTask(task Task) {
	defer func() {
		<-s.workerPool // 释放worker槽位
		s.wg.Done()
	}()

	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	done := make(chan error, 1)
	go func() {
		done <- task.Execute(ctx)
	}()

	select {
	case <-ctx.Done():
		// 处理超时或取消
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			// 记录任务超时
		}
	case err := <-done:
		// 处理任务完成
		if err != nil {
			// 记录任务错误
		}
	}
}

// Submit 提交任务
func (s *Scheduler) Submit(task Task) error {
	select {
	case s.taskQueue <- task:
		return nil
	default:
		return errors.New("task queue is full")
	}
}

// Shutdown 优雅关闭调度器
func (s *Scheduler) Shutdown() {
	close(s.shutdownChan) // 停止接收新任务
	s.wg.Wait()           // 等待所有任务完成
	close(s.taskQueue)    // 关闭任务队列
	close(s.workerPool)   // 关闭worker池
}
