/*
sync1.RWMutex：读写锁，用于在多个读操作和写操作之间提供并发安全性。
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var rwMutex sync.RWMutex
var data int

func readData() {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	fmt.Println("Reading data:", data)
}

func writeData(value int) {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	fmt.Println("Writing data:", value)
	data = value
}

func main() {
	var wg sync.WaitGroup

	// 启动多个读操作
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			readData()
		}()
	}

	// 启动一个写操作
	wg.Add(1)
	go func() {
		defer wg.Done()
		writeData(rand.Intn(100))
	}()

	wg.Wait()
}
