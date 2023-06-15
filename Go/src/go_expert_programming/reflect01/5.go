package main

import (
	"fmt"
	"reflect"
)

// 定义结构体
type User1 struct {
	Id   int
	Name string
	Age  int
}

// 匿名字段
type Boy1 struct {
	User1
	Addr string
}

func main() {
	m := Boy1{User1{1, "zs", 20}, "bj"}
	t := reflect.TypeOf(m)
	fmt.Println(t)

	// Anonymous：匿名
	fmt.Printf("%#v\n", reflect.TypeOf(m).Field(0))
	// 值信息
	fmt.Printf("%#v\n", reflect.ValueOf(m).Field(0))

}
