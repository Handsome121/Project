package main

import (
	"fmt"
	"strconv"
)

func main() {
	//ParseBool()
	//func ParseBool(str string) (value bool, err error)
	//返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误

	//ParseFloat()
	//func ParseFloat(s string, bitSize int) (f float64, err error)
	//解析一个表示浮点数的字符串并返回其值。
	//如果s合乎语法规则，函数会返回最为接近s表示值的一个浮点数（使用IEEE754规范舍入）。
	//bitSize指定了期望的接收类型，32是float32（返回值可以不改变精确值的赋值给float32），64是float64；
	//返回值err是*NumErr类型的，语法有误的，err.Error=ErrSyntax；结果超出表示范围的，返回值f为±Inf，err.Error= ErrRange。

	//ParseInt()
	//func ParseInt(s string, base int, bitSize int) (i int64, err error)
	//返回字符串表示的整数值，接受正负号。
	//base指定进制（2到36），如果base为0，则会从字符串前置判断，”0x”是16进制，”0”是8进制，否则是10进制；
	//bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；
	//返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRange。

	//ParseUint类似ParseInt但不接受正负号，用于无符号整型。
	b, _ := strconv.ParseBool("true")
	i, _ := strconv.ParseInt("-control", 10, 64)
	f, _ := strconv.ParseFloat("3.1415", 64)
	u, _ := strconv.ParseUint("control", 10, 64)
	fmt.Printf("type:%T value:%v\n", b, b)
	fmt.Printf("type:%T value:%v\n", f, f)
	fmt.Printf("type:%T value:%v\n", i, i)
	fmt.Printf("type:%T value:%v\n", u, u)
}
