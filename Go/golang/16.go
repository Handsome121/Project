package main

import "fmt"

//本质是一个指针
type AnimalLF interface {
	Sleep()
	GetColor() string //获取动物的颜色
	GetType() string  //获取动物的种类
}

//具体的类
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}
func (this *Cat) GetColor() string {
	return this.color
}
func (this *Cat) GetType() string {
	return "Cat"
}

//具体的类
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Cat is Sleep")
}
func (this *Dog) GetColor() string {
	return this.color
}
func (this *Dog) GetType() string {
	return "Dog"
}
func showAnimal(animal AnimalLF) {
	animal.Sleep()
	fmt.Println(animal.GetColor())
	fmt.Println(animal.GetType())
}
func main() {
	//var animal AnimalLF //接口的数据类型，父类指针
	//animal = &Cat{"Green"}
	//animal.Sleep()
	cat := Cat{"green"}
	dog := Dog{"yellow"}
	showAnimal(&cat)
	showAnimal(&dog)
}
