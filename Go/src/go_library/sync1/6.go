/*
sync1.Pool：用于实现临时对象池，以提高对象的重用性和性能。
*/
package main

import (
	"fmt"
	"sync"
)

type Object struct {
	// ...
}

func main() {
	pool := &sync.Pool{
		New: func() interface{} {
			return &Object{}
		},
	}

	// 从池中获取对象
	obj := pool.Get().(*Object)
	// 使用对象
	fmt.Println(obj)

	// 将对象放回池中
	pool.Put(obj)

	// 再次从池中获取对象
	obj = pool.Get().(*Object)
	fmt.Println(obj)
}

