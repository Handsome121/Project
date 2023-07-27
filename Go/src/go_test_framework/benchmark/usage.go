/*
用法：
go test -bench -cpu=2 -benchtime=10s/999x -count=2 -benchmem .
结果：
goos: windows	操作系统
goarch: amd64	CPU架构
pkg: benchmark-demo		package名，可以在测试的时候指定package
cpu: 12th Gen Intel(R) Core(TM) i5-1240P	CPU的信息
BenchmarkNewSliceWithCap-8   84  14098503 ns/op  8003589 B/op   1 allocs/op
8表示GOMAXPROCS
84表示b.N的最大值，每次入参b.N的值都在变化
14098503表示循环中每一次的执行时间
8003589表示每次执行fib方法的内存使用量
1平均每次执行fib方法的内存分配次数

指定package：go test -bench benchmark-demo
指定子package：go test -bench benchmark-demo/XXX
当前目录下的所有package：go test -bench ./… （斜杠左侧是一个点，右侧是三个点）

所有以Fib结尾的方法：go test -bench=‘Fib$’ benchmark-demo
所有以BenchmarkNew开始的方法：go test -bench=‘^BenchmarkNew’ benchmark-demo

-cpu调整GOMAXPROCS
-benchtime指定基准测试时长为10秒,999x测试范围就从时间变成了次数
-count用来控制BenchmarkXXX方法的调用次数，而benchtime是用来控制BenchmarkXXX方法的入参b.N的值
-benchmem


*/

package main