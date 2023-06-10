/*
sync1.Cond：条件变量，用于在不同的goroutine之间等待和通知。
*/
package main

import (
	"fmt"
	"sync"
)

var cond sync.Cond
var done bool

func worker(id int) {
	cond.L.Lock()
	for !done {
		cond.Wait()
	}
	cond.L.Unlock()
	fmt.Printf("Worker %d is done\n", id)
}

func main() {
	cond.L = new(sync.Mutex)

	numWorkers := 3

	for i := 1; i <= numWorkers; i++ {
		go worker(i)
	}

	// 做一些其他工作
	// ...

	cond.L.Lock()
	done = true
	cond.Broadcast()
	cond.L.Unlock()

	// 等待所有的worker完成
	// ...

	fmt.Println("All workers are done")
}

