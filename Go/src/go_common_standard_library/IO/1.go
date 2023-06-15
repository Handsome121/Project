package main

import "os"

func main() {
	// 以文件的方式操作终端
	var buf [16]byte
	os.Stdin.Read(buf[:])
	os.Stdin.WriteString(string(buf[:]))
}
