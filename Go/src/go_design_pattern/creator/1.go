/*
单例模式
*/
package main

import (
	"fmt"
	"sync"
)

type SingleData struct {
	data interface{}
}

var data *SingleData
var once sync.Once //通过该类型可以实现单例模式，虽然是多次赋值，但是只执行一次（一个对象多次实例化，但是只有一个，共享对象地址）

func getInstance(i int) *SingleData {
	once.Do(func() {
		data = &SingleData{i}
	})
	return data
}
func main() {
	i1 := getInstance(1)
	fmt.Println(i1)
	i2 := getInstance(2)
	fmt.Println(i2)
	if i1 == i2 {
		fmt.Println("单例模式")
	} else {
		fmt.Println("不是单例模式")
	}
}



