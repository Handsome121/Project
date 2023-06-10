package main

import (
	"fmt"
)

/*
值传递
函数的【形式参数】是对【实际参数】的值拷贝
所有对地址中内容的修改都与外界的实际参数无关
所有基本数据类型作为参数时，都是值传递
*/
func ChangeBaseValue(v int)  {
	fmt.Printf("main: v address=%p\n",&v)
	v *= 2
	//函数返回时所有局部变量（此处包含形式参数v）都被释放了
}

/*
引用传递（地址传递）
函数的【形式参数】记录的是【实际参数】的地址
所有对地址中内容的修改都会直接改变外界的实际参数
*/
func ChangeBaseValuePtr(v *int)  {
	fmt.Printf("main: v address=%p\n",v)
	*v *= 2
}

type Person struct {
	Age int
}

/*
结构体的值传递
形式参数p是对实际参数zhangsan的值拷贝
对p做任何事情都不会对zhangsan产生任何影响
p是函数执行过程中产生的局部变量，函数执行完毕就释放
*/
func ChangeObjectValue(p Person)  {
	fmt.Printf("ChangeObjectValue:p address=%p\n", &p)
	p.Age *= 2
	//函数返回时所有局部变量（此处包含形式参数v）都被释放了
}


/*
结构体的引用传递（地址传递）
形式参数p记录的是实际参数zhansan的内存地址
对p中的内容做任何修改都是对zhangsan的直接修改
p是函数执行过程中产生的局部变量，函数执行完毕就释放
*/
func ChangeObjectValuePtr(p *Person)  {
	fmt.Printf("ChangeObjectValue:p address=%p\n", p)
	p.Age *= 2
	//函数返回时所有局部变量（此处包含形式参数v）都被释放了
}

/*
数组作为形参：值传递
切片和map：现有内容/元素都是引用传递
*/
func ChangeContainerValue(array [3]int, slice []int, mMap map[string]int)  {
	fmt.Printf("ChangeContainerValue:array[0] address=%p\n",&array[0])
	fmt.Printf("ChangeContainerValue:slice[0] address=%p\n",&slice[0])
	fmt.Printf("ChangeContainerValue:map address=%p\n",mMap)

	array[0] = 100
	slice[0] = 200
	mMap["老大"] = 300
}


func ChangeContainerValuePtr(array *[3]int, slice *[]int, mMap *map[string]int)  {
	(*array)[0] = 100
	(*slice)[0] = 200
	(*mMap)["老大"] = 300
	*slice = append(*slice, 4)
	(*mMap)["老四"] = 4
}


/*基本类型的值传递与引用传递*/
func main021() {
	var a = 10
	fmt.Printf("main: a address=%p\n",&a)
	ChangeBaseValue(a)
	fmt.Printf("main: a address=%p\n",&a)
	fmt.Println(a)

	ChangeBaseValuePtr(&a)
	fmt.Printf("main: a address=%p\n",&a)
	fmt.Println(a)
}

/*结构体的值传递与引用传递*/
func main022() {
	zhangsan := Person{Age: 10}
	fmt.Printf("main: zhangsan address=%p\n", &zhangsan)
	//ChangeObjectValue(zhangsan)
	ChangeObjectValuePtr(&zhangsan)
	fmt.Println(zhangsan.Age)
}

func main() {
	var array = [3]int{1,2,3}
	slice := make([]int, 0)
	slice = append(slice, 1, 2, 3)
	mMap := make(map[string]int)
	mMap["老大"] = 1
	mMap["老二"] = 2
	mMap["老三"] = 3
	fmt.Printf("main:array[0] address=%p\n",&array[0])
	fmt.Printf("main:slice[0] address=%p\n",&slice[0])
	fmt.Printf("main:map address=%p\n",mMap)
	ChangeContainerValue(array, slice, mMap)
	//ChangeContainerValuePtr(&array, &slice, &mMap)

	fmt.Println(array)
	fmt.Println(slice)
	fmt.Println(mMap)
}
