package main

import "fmt"

//声明全局变量，方法一二三是可以的
var g0 int
var gA int = 100
var gB = 200

//方法四来声明全局变量
//:=只能用在函数体内来声明
//gC :=200
func main() {
	//方法一：声明一个变量,默认的值是0
	var a int
	fmt.Println("a=", a)
	fmt.Printf("type of a = %T\n", a)

	//方法二：声明一个变量，初始化一个值
	var b int = 100
	var bb string = "abcd"
	fmt.Println("b=", b)
	fmt.Printf("type of b = %T\n", b)
	fmt.Printf("type of bb = %T\n", bb)

	//方法三：初始化的时候可以省略数据类型，通过值自动匹配当前的变量的数据类型
	var c = 100
	var cc = "abcd"
	fmt.Println("c=", c)
	fmt.Printf("type of c = %T\n", c)
	fmt.Printf("type of cc = %T\n", cc)

	//方法四：省去var关键字，直接自动匹配
	e := 200
	fmt.Println("e=", e)
	fmt.Printf("type of f=%T\n", e)

	//声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx=", xx, "yy=", yy)
	var kk, ll = 100, "abcd"
	fmt.Println("kk=", kk, "ll=", ll)
	//多行多变量声明
	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv=", vv, "jj=", jj)
}
