package main

import "fmt"

/*
	    1、明确需要抽象的功能或行为： 首先，确定你需要抽象的功能或行为，这些功能或行为应该适用于多个具体类型。
		例如，你可能需要一个通用的日志记录功能、数据存储功能或网络请求功能。

	    2、定义接口： 在思考清楚需要抽象的功能后，定义一个接口来描述这些功能。接口是一组方法的集合，
		用于定义对象应该具备的行为。在定义接口时，考虑到功能的共性和一致性，而不关注具体的实现细节。

	    3、编写具体类型实现接口： 定义接口后，你可以为不同的具体类型编写实现接口的代码。具体类型需要实现接口中定义的所有方法，
		以满足接口的合约。这些方法的实现将根据具体类型的特点来编写。

	    4、使用接口进行多态操作： 通过接口，你可以以一致的方式操作不同的具体类型对象。这就是多态的概念，
		即不同类型的对象可以被视为同一个接口类型，并且可以通过接口来调用它们的方法。

	    5、依赖接口进行编程： 在编写代码时，应该将依赖关系限制在接口上，而不是具体的实现类上。
		这样可以实现解耦合，提高代码的可测试性和可维护性。使用接口作为参数类型或返回类型，编写与接口相关的代码，而不是与具体类型相关的代码。

	    6、接口的扩展与组合： 接口可以嵌套组合和扩展。你可以通过组合多个接口来创建新的接口，
		从而组合不同的功能。这样可以实现接口的模块化和复用。

	    7、测试接口的实现： 使用接口可以方便地进行单元测试和模拟测试。你可以为具体类型编写针对接口的测试用例，
		确保它们正确实现了接口定义的行为。

	    根据需要调整接口定义： 随着代码的演化和需求的变化，你可能需要对接口进行调整。这可能涉及添加新的方法、调整方法的参数或返回类型等。根据需要进行接口的迭代和演进。
*/
// 定义接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 定义矩形结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// 矩形结构体实现接口中的方法
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 定义圆形结构体
type Circle struct {
	Radius float64
}

// 圆形结构体实现接口中的方法
func (c *Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// 函数接受接口类型作为参数
func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
	fmt.Printf("Perimeter: %f\n", s.Perimeter())
}

func main() {
	rect := Rectangle{Width: 4, Height: 5}
	circle := Circle{Radius: 3}
	PrintShapeInfo(&rect)
	PrintShapeInfo(&circle)
}
