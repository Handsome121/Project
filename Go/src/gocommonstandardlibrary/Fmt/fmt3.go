package main

import "fmt"

func main() {
	var (
		name    string
		age     int
		married bool
	)
	_, _ = fmt.Scanf("1:%s control:%d 3:%t", &name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}
