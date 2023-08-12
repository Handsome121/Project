/*
检查当前Worker队列中是否有可用的Worker，如果有，取出执行当前的task；
没有可用的Worker，判断当前在运行的Worker是否已超过该Pool的容量：{是 —> 再判断工作池是否为非阻塞模式：[是 ——> 直接返回 nil，否 ——> 阻塞等待直至有Worker被放回Pool]，否 —> 新开一个Worker（Goroutine）处理}；
每个Worker执行完任务之后，放回Pool的队列中等待。
*/

package main

import (
	"errors"
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

type sig struct{}

type f func() error

// Pool accept the tasks from client,it limits the total
// of Goroutines to a given number by recycling Goroutines.
type Pool struct {
	// capacity of the pool.
	capacity int32

	// running is the number of the currently running Goroutines.
	running int32

	// expiryDuration set the expired time (second) of every worker.
	expiryDuration time.Duration

	// workers is a slice that store the available workers.
	workers []*Worker

	// release is used to notice the pool to closed itself.
	release chan sig

	// lock for synchronous operation.
	lock sync.Mutex

	once sync.Once

	freeSignal chan sig
}

// Worker is the actual executor who runs the tasks,
// it starts a Goroutine that accepts tasks and
// performs function calls.
type Worker struct {
	// pool who owns this worker.
	pool *Pool

	// task is a job should be done.
	task chan f

	// recycleTime will be updated when putting a worker back into queue.
	recycleTime time.Time
}

// run starts a Goroutine to repeat the process
// that performs the function calls.
func (w *Worker) run() {
	go func() {
		// 循环监听任务列表，一旦有任务立马取出运行
		for f := range w.task {
			if f == nil {
				// 退出Goroutine，运行worker数减一
				w.pool.decRunning()
				return
			}
			_ = f()
			// worker回收复用
			w.pool.putWorker(w)
		}
	}()
}

const DefaultCleanIntervalTime = 10

// NewPool generates an instance of ants pool
func NewPool(size int) (*Pool, error) {
	return NewTimingPool(size, DefaultCleanIntervalTime)
}

var (
	ErrInvalidPoolSize   = errors.New("invalid pool size")
	ErrInvalidPoolExpiry = errors.New("invalid pool expiry")
	ErrPoolClosed        = errors.New("pool is closed")
)

// NewTimingPool generates an instance of ants pool with a custom timed task
func NewTimingPool(size, expiry int) (*Pool, error) {
	if size <= 0 {
		return nil, ErrInvalidPoolSize
	}
	if expiry <= 0 {
		return nil, ErrInvalidPoolExpiry
	}
	p := &Pool{
		capacity:       int32(size),
		freeSignal:     make(chan sig, math.MaxInt32),
		release:        make(chan sig, 1),
		expiryDuration: time.Duration(expiry) * time.Second,
	}
	// 启动定期清理过期worker任务，独立Goroutine运行，
	// 进一步节省系统资源
	go p.monitorAndClear()

	return p, nil
}

// Submit a task to pool
func (p *Pool) Submit(task f) error {
	if len(p.release) > 0 {
		return ErrPoolClosed
	}
	w := p.getWorker()
	w.task <- task

	return nil
}

// getWorker returns an available worker to run the tasks.
func (p *Pool) getWorker() *Worker {
	var w *Worker
	// 标志变量，判断当前正在运行的worker数量是否已到达Pool的容量上限
	waiting := false
	// 加锁，检测队列中是否有可用worker，并进行相应操作
	p.lock.Lock()
	idleWorkers := p.workers
	n := len(idleWorkers)

	// 当前队列中无可用worker
	if n < 1 {
		// 判断运行worker数目已达到该Pool的容量上限，置等待标志
		waiting = p.Running() >= p.Cap()
	} else {
		// 当前队列有可用worker，从队列尾部取出一个使用
		w = idleWorkers[n]
		idleWorkers[n] = nil
		p.workers = idleWorkers[:n]
	}
	// 检测完成，解锁
	p.lock.Unlock()
	// Pool容量已满，新请求等待
	if waiting {
		// 利用锁阻塞等待直到有空闲worker
		for {
			p.lock.Lock()
			idleWorkers = p.workers
			l := len(idleWorkers)

			if l < 1 {
				p.lock.Unlock()
				continue
			}

			w = idleWorkers[l]
			idleWorkers[l] = nil
			p.workers = idleWorkers[:l]
			p.lock.Unlock()

			break
		}
	} else if w == nil {
		// 当前无空闲worker但是Pool还没有满，
		// 则可以直接新开一个worker执行任务
		w = &Worker{
			pool: p,
			task: make(chan f, 1),
		}
		w.run()
		// 运行worker数加一
		p.incRunning()
	}

	return w
}

// putWorker puts a worker back into free pool, recycling the Goroutines.
func (p *Pool) putWorker(worker *Worker) {
	// 写入回收时间，亦即该worker的最后一次结束运行的时间
	worker.recycleTime = time.Now()
	p.lock.Lock()
	p.workers = append(p.workers, worker)
	p.lock.Unlock()
}

// ReSize change the capacity of this pool
func (p *Pool) ReSize(size int32) {
	if size == p.Cap() {
		return
	}
	atomic.StoreInt32(&p.capacity, size)
	diff := int(p.Running() - size)
	if diff > 0 {
		for i := 0; i < diff; i++ {
			p.getWorker().task <- nil
		}
	}
}

// clear expired workers periodically.
func (p *Pool) periodicallyPurge() {
	heartbeat := time.NewTicker(p.expiryDuration)

	for range heartbeat.C {
		currentTime := time.Now()
		p.lock.Lock()
		idleWorkers := p.workers

		if len(idleWorkers) == 0 && p.Running() == 0 && len(p.release) > 0 {
			p.lock.Unlock()
			return
		}
		n := 0

		for i, w := range idleWorkers {
			if currentTime.Sub(w.recycleTime) <= p.expiryDuration {
				break
			}
			n = i
			w.task <- nil
			idleWorkers[i] = nil
		}
		n++

		if n >= len(idleWorkers) {
			p.workers = idleWorkers[:0]
		} else {
			p.workers = idleWorkers[n:]
		}

		p.lock.Unlock()
	}
}

func (p *Pool) decRunning() {
	p.running -= 1
}

func (p *Pool) monitorAndClear() {
	p.periodicallyPurge()
}

func (p *Pool) Running() int32 {
	return p.running
}

func (p *Pool) Cap() int32 {
	return p.capacity
}

func (p *Pool) incRunning() {
	p.running += 1
}

func main() {
	pool, err := NewPool(10)
	if err != nil {
		_ = fmt.Errorf("new pool fail,err is %v", err)
	}
	fmt.Printf("pool is %+v", pool)

	var f1 f = func() error {
		var a = 1
		a += 1
		return nil
	}

	if err = pool.Submit(f1); err != nil {
		_ = fmt.Errorf("submit task fail,err is %v", err)
	}

	var work Worker
	work.run()
}
