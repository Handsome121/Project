package main

import "fmt"

type Human struct {
	name string
	sex  string
}
type SuperMan struct {
	Human //SuperMan类继承了Human类的方法
	level int
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}
func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

//重定义父类的Eat()方法
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

//子类的新方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.fly()...")
}
func main() {
	h := Human{"zhang2", "female"}
	h.Eat()
	h.Walk()
	//定义一个子类对象
	//s := SuperMan{Human{"lisi", "female"}, 10}
	var s SuperMan
	s.name = "li4"
	s.sex = "male"
	s.level = 2
	s.Walk()
	s.Eat()
	s.Fly()
}
