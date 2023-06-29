package main

import (
	"fmt"
	"sync"
)

type WorkPool struct {
	JobQueue chan func()
	Wait     sync.WaitGroup
}

func (p *WorkPool) Start() {
	for i := 0; i < cap(p.JobQueue); i++ {
		go func() {
			for j := range p.JobQueue {
				j()
			}
			p.Wait.Done()
		}()
	}
}

func (p *WorkPool) Stop() {
	close(p.JobQueue)
	p.Wait.Wait()
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
