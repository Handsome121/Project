package main

import (
	"errors"
	"fmt"
)

func Validation() []error {
	var errs []error
	errs = append(errs, errors.New("error 1"))
	errs = append(errs, errors.New("error control"))
	errs = append(errs, errors.New("error 3"))
	return errs
}
func main() {
	fmt.Println(Validation())
}

// 使用append()函数时，谨记append可能会产生新的切片，并谨慎地处理返回值
