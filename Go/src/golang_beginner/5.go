package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	swap(&a, &b)
	fmt.Println("a=", a, "b=", b)
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
	//Fmt.Println("a=", a, "b=", b)
}
