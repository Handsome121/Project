package main

import (
	"fmt"
	"reflect"
)
//=======================很关键========================
//什么是反射？ 反射是指在程序运行期对程序本身进行访问和修改的能力
//变量的内在机制：
//变量包含类型信息和值信息 var arr [10]int arr[0] = 10
//类型信息：是静态的元信息，是预先定义好的
//值信息：是程序运行过程中动态改变的

func reflectType(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println("类型是:", t)
	// kind()可以获取具体类型
	k := t.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Printf("a is float64\n")
	case reflect.String:
		fmt.Printf("string")
	}
}

func main() {
	var x float64 = 3.4
	reflectType(x)
}
