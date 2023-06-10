package main

import (
	"fmt"
	"strconv"
)

/*Format系列函数实现了将给定类型数据格式化为string类型数据的功能*/

func main() {
	//FormatBool()
	//func FormatBool(b bool) string
	//根据b的值返回”true”或”false”。

	//FormatInt()
	//func FormatInt(i int64, base int) string
	//返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母’a’到’z’表示大于10的数字。

	//FormatUint()
	//是FormatInt的无符号整数版本。

	//FormatFloat()
	//func FormatFloat(f float64, fmt byte, prec, bitSize int) string
	//函数将浮点数表示为字符串并返回。
	//bitSize表示f的来源类型（32：float32、64：float64），会据此进行舍入。
	//fmt表示格式：’f’（-ddd.dddd）、’b’（-ddddp±ddd，指数为二进制）、’e’（-d.dddde±dd，十进制指数）、’E’（-d.ddddE±dd，十进制指数）、
	//’g’（指数很大时用’e’格式，否则’f’格式）、’G’（指数很大时用’E’格式，否则’f’格式）。
	//prec控制精度（排除指数部分）：对’f’、’e’、’E’，它表示小数点后的数字个数；对’g’、’G’，它控制总的数字个数。如果prec 为-1，则代表使用最
	//少数量的、但又必需的数字来表示f。
	s1 := strconv.FormatBool(true)
	s2 := strconv.FormatFloat(1234567891234567890, 'f', -1, 64)
	s3 := strconv.FormatInt(15, 16)
	s4 := strconv.FormatUint(2, 16)
	fmt.Printf("type:%T value:%v\n", s1, s1)
	fmt.Printf("type:%T value:%v\n", s2, s2)
	fmt.Printf("type:%T value:%v\n", s3, s3)
	fmt.Printf("type:%T value:%v\n", s4, s4)
}

