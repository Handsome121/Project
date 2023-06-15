/*
sync1.Map：并发安全的字典结构，可用于多个goroutine之间的安全访问。
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	m.Store("key1", "value1")
	m.Store("key2", "value2")

	value, ok := m.Load("key1")
	if ok {
		fmt.Println("Value of key1:", value)
	}

	m.Delete("key2")

	m.Range(func(key, value interface{}) bool {
		fmt.Println("Key:", key, "Value:", value)
		return true
	})
}
