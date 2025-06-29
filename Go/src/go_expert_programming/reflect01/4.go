package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("hello")
}

func Point(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("类型", t)
	fmt.Println("字符串类型", t.Name())
	// 获取值
	v := reflect.ValueOf(o)
	fmt.Println(v)
	// 获取所有属性
	// 获取结构体字段个数：t.NumFiled()
	for i := 0; i < t.NumField(); i++ {
		// 取每个字段
		f := t.Field(i)
		fmt.Printf("%s:%v\n", f.Name, f.Type)
		// 获取字段的值信息
		// interface():获取字段对应的值
		val := v.Field(i).Interface()
		fmt.Println("val", val)
	}
	fmt.Println("===============方法===============")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name)
		fmt.Println(m.Type)

	}
}

func main() {
	u := User{1, "zs", 20}
	Point(u)
}
