package main

import (
	"testing"
	"time"
)

// BenchmarkFibWithPrepare 进入正式测试前需要耗时做准备工作的case
func BenchmarkFibWithPrepare(b *testing.B) {
	// 假设这里有个耗时800毫秒的初始化操作
	<-time.After(800 * time.Millisecond)

	// 这样表示重新开始计时,前面的耗时与本次基准测试无关
	//b.ResetTimer()

	// 这下面才是咱们真正想做基准测试的代码
	for n := 0; n < b.N; n++ {
		b.StartTimer()
		fib(30)
		b.StopTimer()
	}

	// 假设这里有个耗时10毫秒的结束操作
	<-time.After(10 * time.Millisecond)

}


