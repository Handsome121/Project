package main

import "fmt"

type AInterface interface {
	Say()
}
type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("Stu Say()...")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer Say i =", i)
}

type BInterface interface {
	Hello()
}
type Monster struct {
}

func (m Monster) Hello() {
	fmt.Println("Monster Hello()...")
}
func (m Monster) Say() {
	fmt.Println("Monster Say()...")
}
func main() {
	var stu Stu
	var a AInterface = stu
	a.Say()

	var i integer = 10
	var b AInterface = i
	b.Say()

	var monster Monster
	var a2 AInterface = monster
	var b2 BInterface = monster
	a2.Say()
	b2.Hello()
}

//golang接口实现注意点:
//1、接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量（实例）
//2、接口中所有的方法都没有方法体，即都是没有实现的方法
//3、在golang中，一个自定义类型需要将某个接口的所有方法都实现，我们说这个自定义类型是实现了该接口
//4、只要是系定义类型，就可以实现接口，而不仅仅是结构体
//5、一个自定义类型可以实现多个接口
//6、golang接口中不能有任何变量
//7、一个接口(A)可以继承多个别(B,C)的接口，这时如果要实现A接口，也必须将B,C接口的方法也全部实现，并且B,C中不能有重名方法
//8、interface类型默认是一个指针(引用类型),如果没有对interface初始化就使用那么会输出nil
//9、空接口interface{}没有任何方法,那么所有的类型都实现了空接口
//10、注意实现接口的是实例还是指针
