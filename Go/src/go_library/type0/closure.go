package main

import "fmt"

type Counter struct {
	count int
}

type CounterFunc func() int

func (c *Counter) GetCount() int {
	return c.count
}

func (c *Counter) Increment() {
	c.count++
}

func NewCounter() CounterFunc {
	counter := &Counter{}

	incrementFunc := func() int {
		counter.Increment()
		return counter.GetCount()
	}

	return incrementFunc
}

func main() {
	counter := NewCounter()

	fmt.Println(counter()) // 输出 "1"
	fmt.Println(counter()) // 输出 "2"
	fmt.Println(counter()) // 输出 "3"
}
