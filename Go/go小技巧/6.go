package main

import "time"

func main() {
	//用 time.Duration 代替 int64 + 变量名
	// BAD
	var delayMillis int64 = 15000

	// GOOD
	var delay time.Duration = 15 * time.Second
}
