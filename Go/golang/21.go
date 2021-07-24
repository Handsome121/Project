package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) called() {
	fmt.Println("user is called...")
	fmt.Printf("%v\n", this)
}
func DoFieldAndMethod(input interface{}) {
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is :", inputType.Name())
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is :", inputValue)
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s:%v=%v\n", field.Name, field.Type, value)
	}
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s:%v\n", m.Name, m.Type)
	}

}
func main() {
	user := User{1, "zhang2", 3}
	//user.called()
	DoFieldAndMethod(user)
}
