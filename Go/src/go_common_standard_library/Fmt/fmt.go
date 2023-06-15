package main

import "fmt"

//func main() {
//	Fmt.Print("在终端打印该信息。")
//	name := "枯藤"
//	Fmt.Printf("我是：%s\n", name)
//	Fmt.Println("在终端打印单独一行显示")
//}

//func main() {
//	// 向标准输出写入内容
//	_, _ = Fmt.Fprintln(os.Stdout, "向标准输出写入内容")
//	fileObj, err := os.OpenFile("./123.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
//	if err != nil {
//		Fmt.Println("打开文件出错，err:", err)
//		return
//	}
//	name := "枯藤"
//	// 向打开的文件句柄中写入内容
//	_, _ = Fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)
//}

//func main() {
//	//s1 := Fmt.Sprint("枯藤")
//	//name := "枯藤"
//	//age := 18
//	//s2 := Fmt.Sprintf("name:%s,age:%d", name, age)
//	//s3 := Fmt.Sprintln("枯藤")
//	//Fmt.Println(s1, s2, s3)
//
//	_ = Fmt.Errorf("这是一个错误")
//}

func main() {
	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)
	o := struct{ name string }{"枯藤"}
	fmt.Printf("%v\n", o)
	fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)
	fmt.Printf("100%%\n")
	//%v	值的默认格式表示
	//%+v	类似%v，但输出结构体时会添加字段名
	//%#v	值的Go语法表示
	//%T	打印值的类型
	//%%	百分号

	n := 65
	fmt.Printf("%b\n", n)
	fmt.Printf("%c\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X\n", n)
	//%b	表示为二进制
	//%c	该值对应的unicode码值
	//%d	表示为十进制
	//%o	表示为八进制
	//%x	表示为十六进制，使用a-f
	//%X	表示为十六进制，使用A-F
	//%U	表示为Unicode格式：U+1234，等价于”U+%04X”
	//%q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示

	f := 12.34
	fmt.Printf("%b\n", f)
	fmt.Printf("%e\n", f)
	fmt.Printf("%E\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%g\n", f)
	fmt.Printf("%G\n", f)
	//%b	无小数部分、二进制指数的科学计数法，如-123456p-78
	//%e	科学计数法，如-1234.456e+78
	//%E	科学计数法，如-1234.456E+78
	//%f	有小数部分但无指数部分，如123.456
	//%F	等价于%f
	//%g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
	//%G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）

	s := "枯藤"
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("%X\n", s)
	//%s	直接输出字符串或者[]byte
	//%q	该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
	//%x	每个字节用两字符十六进制数表示（使用a-f
	//%X	每个字节用两字符十六进制数表示（使用A-F）

	a := 18
	fmt.Printf("%p\n", &a)
	fmt.Printf("%#p\n", &a)
	//%p	表示为十六进制，并加上前导的0x

	c := 88.88
	fmt.Printf("%f\n", c)
	fmt.Printf("%9f\n", c)
	fmt.Printf("%.2f\n", c)
	fmt.Printf("%9.2f\n", c)
	fmt.Printf("%9.f\n", c)
	//%f	默认宽度，默认精度
	//%9f	宽度9，默认精度
	//%.2f	默认宽度，精度2
	//%9.2f	宽度9，精度2
	//%9.f	宽度9，精度0fmt3
}
