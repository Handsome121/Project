package main

import "fmt"

//定义一个类
type Hero struct {
	Name  string
	Ad    int
	Level int
}

/*func (this Hero) Show() {
	Fmt.Println("Name = ", this.Name)
	Fmt.Println("Ad = ", this.Ad)
	Fmt.Println("Level = ", this.Level)
}
func (this Hero) GetName() string {
	return this.Name
}
func (this Hero) setName(newName string) {
	//this 是调用该方法的对象的一个副本（拷贝）
	this.Name = newName
}*/

func (this *Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}
func (this *Hero) GetName() string {
	return this.Name
}
func (this *Hero) setName(newName string) {
	//this 是调用该方法的对象的一个副本（拷贝）
	this.Name = newName
}
func main() {
	hero := Hero{Name: "zhang2", Ad: 100, Level: 1}
	hero.Show()
	hero.setName("lisi")
	hero.Show()
}
