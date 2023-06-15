package main

import (
	"fmt"
	"strconv"
	"strings"
)

//字符串常用的系统函数
func main() {
	//统计字符串的长度，按字节len(str)
	//golang的编码统一为utf-8,占一个字节，汉字占三个字节

	str := "hello北"
	fmt.Println("str len=", len(str))
	str2 := "hello北京"
	//字符串遍历
	//r := []rune(str2)
	//for i := 0; i < len(r); i++ {
	//	Fmt.Printf("字符=%c\n", r[i])
	//}
	for _, value := range str2 {
		fmt.Printf("字符=%c\n", value)
	}

	//字符串转整数
	n, err := strconv.Atoi("control")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换的结果是", n)
	}

	//整数转字符串
	str = strconv.Itoa(123456)
	fmt.Printf("str=%v,str=%T", str, str)

	//字符串转[]byte类型
	var bytes = []byte("hello go")
	fmt.Printf("byte=%v\n", bytes)

	//[]byte转字符串
	str = string([]byte{97, 98, 99})
	fmt.Printf("str=%v\n", str)

	//10进制转2，8，16，进制
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的2进制=%v\n", str)

	str = strconv.FormatInt(123, 8)
	fmt.Printf("123对应的8进制=%v\n", str)

	str = strconv.FormatInt(123, 16)
	fmt.Printf("123对应的16进制=%v\n", str)

	//判断字串是否包含在字符串中
	b := strings.Contains("seafood", "mary")
	fmt.Printf("b=%v\n", b)

	//返回字符串中有几个指定的字串
	num := strings.Count("ce", "e")
	fmt.Printf("num=%v\n", num)

	//不区分大小写的字符串比较
	b = strings.EqualFold("abc", "Abc")
	fmt.Printf("b=%v\n", b)

	fmt.Println("结果", "abc" == "Abc")

	//返回字串在字符串中第一次出现的index值
	index := strings.Index("abc", "b")
	fmt.Printf("index=%v\n", index)

	//返回字串在字符串中最后一次出现的index值
	index2 := strings.LastIndex("abbbbbc", "b")
	fmt.Printf("index=%v\n", index2)

	//替换指定字符串
	str2 = "go go hello"
	str = strings.Replace(str2, "go", "北京", -1)
	fmt.Printf("str=%v\n", str)
	fmt.Printf("str2=%v\n", str2)

	//按照指定的某个字符，将一个字符串分割成字符串数组
	strArr := strings.Split("hello,world,ok", ",")
	fmt.Printf("strArr=%v\n", strArr)

	//字符串转大小写
	c := strings.ToLower("ABC")
	fmt.Printf("c=%v\n", c)

	d := strings.ToUpper("abc")
	fmt.Printf("d=%v\n", d)

	//去除字符串左右两边的空格
	e := strings.TrimSpace(" 你好   ")
	fmt.Printf("e=%v\n", e)

	//去除字符串左右两边指定的字符去掉
	f := strings.Trim("!hello !", "!")
	fmt.Printf("f=%v\n", f)

	//去除字符串左边指定的字符去掉
	g := strings.TrimLeft("! hello !", " !")
	fmt.Printf("g=%v\n", g)

	//去除字符串左边指定的字符去掉
	h := strings.TrimRight("! hello !", " !")
	fmt.Printf("h=%v\n", h)

	//判断字符串是否以指定的字符串开头
	i := strings.HasPrefix("ftp://192.168.10.1", "hsp")
	fmt.Printf("i=%v\n", i)
	//判断字符串是否以指定的字符串结尾
	k := strings.HasSuffix("ftp://192.168.10.1", "hsp")
	fmt.Printf("k=%v\n", k)
}
