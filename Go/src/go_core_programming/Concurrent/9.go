package main

import (
	"fmt"
	"sync"
)

type WorkPool struct {
	JobQueue chan func()
	StopChan chan struct{}
	Wait     sync.WaitGroup
}

func (p *WorkPool) Start() {
	for i := 0; i < cap(p.JobQueue); i++ {
		go func() {
			for {
				select {
				case job, ok := <-p.JobQueue:
					if !ok {
						// 任务队列已关闭，退出循环
						return
					}
					job()
					p.Wait.Done()
				case <-p.StopChan:
					// 收到停止信号，退出循环
					return
				}
			}
		}()
	}
}

func (p *WorkPool) Stop() {
	close(p.JobQueue)
	p.Wait.Wait()
	close(p.StopChan)
}

func (p *WorkPool) Submit() {
	p.Wait.Add(1)
	p.JobQueue <- func() {
		fmt.Println("123")
	}
}

func main() {
	pool := WorkPool{
		JobQueue: make(chan func(), 5),
		Wait:     sync.WaitGroup{},
	}
	pool.Start()
	pool.Submit()

}
