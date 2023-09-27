/*
装饰器模式
*/
package main

import "fmt"

type Component interface {
	Calc() int
}
type ConcreteComponent struct{}

func (*ConcreteComponent) Calc() int {
	return 0
}

// MulDecorator 乘法修饰器
type MulDecorator struct {
	Component
	num int
}

func (d *MulDecorator) Calc() int {
	//return (d.Component.Calc()+1)*d.num
	return d.Component.Calc() * d.num
}
func NewMulDecorator(c Component, data int) *MulDecorator {
	return &MulDecorator{
		Component: c,
		num:       data,
	}
}

// AddDecorator 加法修饰器
type AddDecorator struct {
	Component
	num int
}

func (d *AddDecorator) Calc() int {
	return d.Component.Calc() + d.num
}
func NewAddDecorator(c Component, data int) *AddDecorator {
	return &AddDecorator{
		Component: c,
		num:       data,
	}
}
func main() {
	var c Component
	//指针要先进行创建（初始化），不然直接赋值会出现空指针异常，即：给指针为nil的对象赋值
	//错误提示：panic: runtime error: invalid memory address or nil pointer dereference
	c = &ConcreteComponent{}
	c = NewAddDecorator(c, 10) //先调用加
	c = NewMulDecorator(c, 8)  //在以上基础之上再执行乘法
	fmt.Println("加计算结果：", c.Calc())

}
