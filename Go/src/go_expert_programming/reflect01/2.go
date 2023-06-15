package main

import (
	"fmt"
	"reflect"
)

//反射三定律
func main() {
	////反射第一定律:反射可以将interface类型变量转换成反射对象
	//var x float64 = 3.4
	//t := reflect.TypeOf(x)
	//fmt.Println("type:", t)
	//
	//v := reflect.ValueOf(x)
	//fmt.Println("value:", v)
	////tips:
	////变量x在传入reflect.TypeOf()函数时，实际上做了一次转换，x变量被转换成一个空接口传入，reflect.ValueOf也是如此

	//反射第二定律:反射可以将反射对象还原成interface类型变量
	//var a interface{}
	//a = 100
	//fmt.Printf("类型是%T\n", a)
	//v := reflect.ValueOf(a)
	//b := v.Interface()
	//fmt.Printf("类型是%T\n", b)
	//if a == b {
	//	fmt.Printf("they are same!\n")
	//}

	//反射第三定律:反射对象可修改，value值必须是可设置的
	var x float64 = 3.4
	v := reflect.ValueOf(&x)
	v.Elem().SetFloat(7.1)
	fmt.Println("x:", v.Elem().Interface())
}
