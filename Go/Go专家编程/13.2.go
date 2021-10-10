package main

import (
	"errors"
	"fmt"
)

func ValidateName(name string) error {
	if name != "" {
		return nil
	}
	return errors.New("empty name")
}
func Validations(name string) []error {
	var errs []error
	errs = append(errs, ValidateName(name))
	return errs
}
func main() {
	fmt.Println(Validations(""))
}

// 使用append()函数时，append会追加nil值，应该尽量避免追加无意义的元素
