package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	//flag.Type()
	//name := flag.String("name", "张三", "姓名")
	//age := flag.Int("age", 18, "年龄")
	//married := flag.Bool("married", false, "婚否")
	//delay := flag.Duration("d", 0, "时间间隔")
	//fmt.Println(name)
	//fmt.Println(age)
	//fmt.Println(married)
	//fmt.Println(delay)

	//flag.TypeVar()
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "时间间隔")
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(married)
	fmt.Println(delay)

	//flag.Parse()
	//通过以上两种方法定义好命令行flag参数后，需要通过调用flag.Parse()来对命令行参数进行解析。
	//
	//支持的命令行参数格式有以下几种：
	//
	//-flag xxx （使用空格，一个-符号）
	//--flag xxx （使用空格，两个-符号）
	//-flag=xxx （使用等号，一个-符号）
	//--flag=xxx （使用等号，两个-符号）
	//其中，布尔类型的参数必须使用等号的方式指定。
	//
	//Flag解析在第一个非flag参数（单个”-“不是flag参数）之前停止，或者在终止符”–“之后停止。

	//其他
	//flag.Args()  //返回命令行参数后的其他参数，以[]string类型
	//flag.NArg()  //返回命令行参数后的其他参数个数
	//flag.NFlag() //返回使用的命令行参数个数

}
