package main

import "fmt"

func main() {
	var (
		name    string
		age     int
		married bool
	)
	_, _ = fmt.Scanln(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

	//Fmt.Scanln遇到回车就结束扫描了，这个比较常用。
}
