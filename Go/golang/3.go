package main

import "fmt"

//comst来定义枚举类型
const (
	//每行的iota都会累加1,第一行的iota的默认值是0
	BEIJING  = iota //iota=0
	SHANGHAI        //iota=1
	SHENZHEN        //iota=2
)
const (
	a, b = iota + 1, iota + 2 //iota=0
	c, d                      //iota=1
	e, f                      //iota=2
	g, h = iota * 2, iota * 3 //iota=3
	i, k                      //iota=4
)

//iota只能够配合const()一起使用,iota只有在const进行累加效果
func main() {
	//常量
	const length int = 10
	fmt.Println("length=", length)
	fmt.Println("BEIJING=", BEIJING)
	fmt.Println("BEIJING=", SHANGHAI)
	fmt.Println("BEIJING=", SHENZHEN)
}
