package main

// 斐波拉契数列
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return fib(n-2) + fib(n-1)
}
