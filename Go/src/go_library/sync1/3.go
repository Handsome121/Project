/*
sync1.WaitGroup：等待一组goroutine完成执行。
*/
package main

import (
	"sync"
	"fmt"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	// 执行具体的工作
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	numWorkers := 3
	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers done")
}
