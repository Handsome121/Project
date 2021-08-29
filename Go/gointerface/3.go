package main

import "fmt"

type Monkey struct {
	Name string
}
type BirdAble interface {
	Flying()
}

func (m *Monkey) climbing() {
	fmt.Println(m.Name, "生来会爬树")
}

type LittleMonkey struct {
	Monkey // 继承
}

func (p *LittleMonkey) Flying() {
	fmt.Println(p.Name, "通过学习会飞翔")
}
func main() {
	//创建一个LittleMonkey实例
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	monkey.climbing()
	monkey.Flying()
}

//接口与继承实现注意点：
//1、继承的价值在于：解决代码的复用性和可维护性;接口的价值主要在于:设计好各种规范，让其他自定义类型去实现这些方法
//2、接口比继承更加灵活，继承是满足is-a的关系，而接口只满足like-a的关系
//3、接口在一定程度上实现代码解耦
