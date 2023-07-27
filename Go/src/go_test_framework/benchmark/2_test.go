package main

import "testing"

const (
	SliceLengthMillion        = 1000000   // 往切片中添加数据的长度，百万
	SliceLengthTenMillion     = 10000000  // 往切片中添加数据的长度，千万
	SliceLengthHundredMillion = 100000000 // 往切片中添加数据的长度，亿
)

func BenchmarkNewSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		newSlice(SliceLengthMillion)
	}
}

func BenchmarkNewSliceWithCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		newSliceWithCap(SliceLengthMillion)
	}
}
