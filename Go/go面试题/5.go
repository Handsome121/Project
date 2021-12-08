package main

//1.通过指针变量 p 访问其成员变量 name，有哪几种方式？
//A.p.name
//B.(&p).name
//C.(*p).name
//D.p->name
//参考答案及解析：AC。& 取址运算符，* 指针解引用。

//2.下面这段代码能否通过编译？如果通过，输出什么？
//package main
//
//import "fmt"
//
//type MyInt1 int  基于类型int创建了新类型 MyInt1
//type MyInt2 = int   创建了int的类型别名 MyInt2
//
//func main() {
//	var i int =0
//	var i1 MyInt1 = i  将int类型的变量赋值给MyInt1类型的变量
//	var i2 MyInt2 = i  而MyInt2只是int的别名，本质上还是int，可以赋值。
//	fmt.Println(i1,i2)
//}

