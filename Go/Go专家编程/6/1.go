package main

import (
	"fmt"
	"reflect"
)

//反射定律:
//1、第一定律：反射可以将interface类型变量转换成反射对象
//func main() {
//	var x float64 = 3.4
//	t := reflect.TypeOf(x)
//	fmt.Println(t)
//	v := reflect.ValueOf(x)
//	fmt.Println(v)
//}

//2、第二定律：反射可以将反射对象还原成interface对象
//func foo() {
//	var A interface{}
//	A = 100
//	v := reflect.ValueOf(A) //获取接口变量A的反射对象
//	t := reflect.TypeOf(A)
//	fmt.Println(v)
//	fmt.Println(t)
//	B := v.Interface() //通过反射对象的interface()获取B
//	fmt.Println(B)
//	if A == B {
//		fmt.Printf("they are same!\n")
//	}
//}
//func main() {
//	foo()
//}

//3、反射对象可修改，value值必须是可设置的
func main() {
	//var x float64 = 3.4
	//v := reflect.ValueOf(x)
	//v.SetFloat(7.1) //触发panic

	var x float64 = 3.4
	v := reflect.ValueOf(&x)
	v.Elem().SetFloat(7.1)
	fmt.Println("x:", v.Elem().Interface())
}
