package main

import "fmt"

func foo() {
	var out []*int
	for i := 0; i < 3; i++ {
		iCopy := i
		out = append(out, &iCopy)
	}
	fmt.Print("values:", *out[0], *out[1], *out[2])
}
func main() {
	foo()
}
