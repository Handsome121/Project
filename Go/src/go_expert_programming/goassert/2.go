/*类型断言的最佳实践*/
package main

import "fmt"

type Student struct {
}

func TypeJudge(items ...interface{}) {
	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是bool类型，值是%v\n", index, x)
		case float32:
			fmt.Printf("第%v个参数是bool类型，值是%v\n", index, x)
		case int, int32, int64:
			fmt.Printf("第%v个参数是整数类型，值是%v\n", index, x)
		case string:
			fmt.Printf("第%v个参数是string类型，值是%v\n", index, x)
		case Student:
			fmt.Printf("第%v个参数是Student类型，值是%v\n", index, x)
		case *Student:
			fmt.Printf("第%v个参数是*Student类型，值是%v\n", index, x)
		default:
			fmt.Printf("输入参数不合法")
		}
	}
}

func main() {
	var a string = "hello"
	var b int = 123
	var c bool = true
	stu1 := Student{}
	stu2 := &Student{}
	TypeJudge(a, b, c, stu1, stu2)
}
