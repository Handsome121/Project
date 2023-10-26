package main

import "fmt"

type MyStruct struct {
	NestedField []*NestedStruct
}

type NestedStruct struct {
	Name  string
	Value int
}

func main() {
	// 初始化嵌套切片类型的结构体
	myStruct := MyStruct{
		NestedField: []*NestedStruct{
			{Name: "A", Value: 1},
			{Name: "B", Value: 2},
			{Name: "C", Value: 3},
		},
	}

	// 遍历切片并打印每个指针所指的值
	for _, nested := range myStruct.NestedField {
		fmt.Println(*nested)
	}
}
