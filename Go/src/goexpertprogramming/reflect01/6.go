package main

import (
	"fmt"
	"reflect"
)

// 定义结构体
type User2 struct {
	Id   int
	Name string
	Age  int
}

// 修改结构体值
func SetValue(o interface{}) {
	v := reflect.ValueOf(o)
	// 获取指针指向的元素
	v = v.Elem()
	// 取字段
	f := v.FieldByName("Name")
	if f.Kind() == reflect.String {
		f.SetString("kuteng")

	}
}
func main() {
	u := User2{1, "5lmh.com", 20}
	SetValue(&u)
	fmt.Println(u)
}
