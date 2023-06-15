package main

import (
	"fmt"
	"reflect"
)

func reflectValue(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Println(v)
	k := v.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Println("a是:", v.Float())
	}
}

func main() {
	var x float64 = 3.4
	reflectValue(x)
}
