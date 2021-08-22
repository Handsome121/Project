package main

import (
	"fmt"
	"strconv"
	"time"
)

func test03() {
	str := ""
	for i := 0; i < 1000000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}
func main() {
	//在执行test03前先获取到unix时间戳
	start := time.Now().Unix()
	fmt.Printf("开始时间为%v\n", start)
	test03()
	end := time.Now().Unix()
	fmt.Printf("结束时间为%v\n", end)
	fmt.Printf("执行时间为%v\n", end-start)
}
