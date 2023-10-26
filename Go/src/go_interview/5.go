package main

import "fmt"

func main() {
	var s = "1234"
	for _, value := range []rune(s) {
		fmt.Println(string(value))
	}
}
