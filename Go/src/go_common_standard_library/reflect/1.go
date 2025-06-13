/*
1. 性能考虑：反射操作比直接代码慢很多，避免在性能关键路径使用
2. 安全性：确保类型检查，避免运行时panic
3. 可维护性：反射代码通常较难理解和维护，考虑是否有必要使用反射
*/
package main

import (
	"fmt"
	"reflect"
)

// 基本用法
func inspect(x interface{}) {
	t := reflect.TypeOf(x)  // 获取类型信息
	v := reflect.ValueOf(x) // 获取值信息

	fmt.Println("Type:", t)
	fmt.Println("Value:", v)
}

// 获取类型信息
func showTypeInfo(x interface{}) {
	t := reflect.TypeOf(x)

	fmt.Println("Kind:", t.Kind())           // 底层类型(Int, Struct, Ptr等)
	fmt.Println("Name:", t.Name())           // 类型名称
	fmt.Println("Size:", t.Size())           // 类型大小(字节)
	fmt.Println("NumMethod:", t.NumMethod()) // 方法数量
}

func inspectStruct(x interface{}) {
	t := reflect.TypeOf(x)

	// 遍历结构体字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field %d: %s (type %v, tag %v)\n",
			i, field.Name, field.Type, field.Tag)
	}
}

type Person struct {
	Name string
	Age  int
}

/*
关键点说明
1. 必须使用指针：
reflect.ValueOf(&struct) 而不是 reflect.ValueOf(struct)
否则会报错"unaddressable value"
2. Elem() 方法：
用于解引用指针，获取指针指向的实际值
3. 字段检查：
IsValid() 检查字段是否存在
CanSet() 检查字段是否可设置
Kind() 检查字段类型
4. 设置不同类型的值：
SetString() 设置字符串
SetInt() 设置整数
SetFloat() 设置浮点数
*/
func modifyStructValue(p interface{}) {
	// 1. 获取指针的Value
	v := reflect.ValueOf(p)

	// 2. 检查是否是指针类型
	if v.Kind() != reflect.Ptr {
		fmt.Println("必须传入结构体指针")
		return
	}

	// 3. 获取指针指向的实际结构体
	v = v.Elem()

	// 4. 检查是否是结构体类型
	if v.Kind() != reflect.Struct {
		fmt.Println("传入的不是结构体指针")
		return
	}

	// 5. 修改Name字段
	nameField := v.FieldByName("Name")
	if nameField.IsValid() && nameField.CanSet() {
		if nameField.Kind() == reflect.String {
			nameField.SetString("Modified Name")
		}
	}

	// 6. 修改Age字段
	ageField := v.FieldByName("Age")
	if ageField.IsValid() && ageField.CanSet() {
		if ageField.Kind() == reflect.Int {
			ageField.SetInt(30)
		}
	}
}

type Calculator struct{}

func (c Calculator) Add(a, b int) int {
	return a + b
}

func callMethod(x interface{}, methodName string, args ...interface{}) {
	v := reflect.ValueOf(x)
	method := v.MethodByName(methodName)

	// 准备参数
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}

	// 调用方法
	out := method.Call(in)

	// 处理返回值
	for _, val := range out {
		fmt.Println(val.Interface())
	}
}

func main() {
	//inspect(21)
	//inspect("hello")
	//inspect(true)
	//showTypeInfo(21)
	//showTypeInfo("hello")
	//showTypeInfo(true)
	//type Person struct {
	//	Name string `json:"name"`
	//	Age  int    `json:"age"`
	//}
	//showTypeInfo(test{Name: "test", Age: 18})
	//inspectStruct(Person{Name: "test", Age: 18})

}
