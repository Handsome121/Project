package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	var b Point
	b = a.(Point) //类型断言 表示a是否指向Point类型的变量，如果是就转成Point类型并赋给b变量，否则报错
	fmt.Println(b)

	//var x interface{}
	//var b2 float32 = 1.1
	//x = b2 //空接口，可以接受任意类型
	//y := x.(float32)
	//Fmt.Printf("y的类型是%T,值是%v\n", y, y)

	/*带检测的类型断言*/
	var x interface{}
	var b2 float64 = 1.1
	x = b2 //空接口，可以接受任意类型
	if y, ok := x.(float64); ok {
		fmt.Println("转换成功")
		fmt.Printf("y的类型是%T,值是%v\n", y, y)
	} else {
		fmt.Println("转换失败")
	}

	//主要说明:
	//在进行类型断言时，如果类型不匹配，就会报panic，因此进行类型断言时，要确保原来空接口指向的就是对应的类型
	//需要带上检测机制
}
