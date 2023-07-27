package main

import "testing"

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(30)
	}
}

// 并行测试
func BenchmarkParallelFib(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fib(30)
		}
	})
}

