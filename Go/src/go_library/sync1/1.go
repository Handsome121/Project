/*
sync1.Mutex：互斥锁，用于保护共享资源的访问。
*/
package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var counter int

func increment() {
	mutex.Lock()
	counter++
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 5
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()

	fmt.Println("Counter:", counter)
}
