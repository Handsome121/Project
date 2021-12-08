package main

import "time"

func main() {
	//30 * time.Second 比 time.Duration(30) * time.Second 更好

	//你不需要将无类型的常量包装成类型，编译器会找出来。
	//另外最好将常量移到第一位：
	// BAD
	delay := time.Second * 60 * 24 * 60

	// VERY BAD
	delay := 60 * time.Second * 60 * 24

	// GOOD
	delay := 24 * 60 * 60 * time.Second

}