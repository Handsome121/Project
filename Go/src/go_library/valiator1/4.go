///*
//验证map,slice
//*/
//
//package main
//
//import (
//	"fmt"
//
//	"github.com/go-playground/validator/v10"
//)
//
//type User struct {
//	// 第一个 required 规则用于确保切片本身不为零值（即不为 nil）。
//	// dive 规则指定了对切片元素进行递归验证。
//	// 第二个 required 规则用于验证切片中的每个元素是否不为空。
//	Accounts []string `validate:"required,dive,required"`
//
//	// 第一个 required 规则用于确保 map 本身不为零值（即不为 nil）。
//	// dive 规则指定了对 map 值进行递归验证。
//	// gt=0 规则用于验证 map 中的每个值是否大于 0。
//	// 第二个 required 规则用于验证 map 中的每个值是否不为空。
//	Balance map[string]int `validate:"required,dive,gt=0,required"`
//}
//
//var validate *validator.Validate
//
//func main() {
//	validate = validator.New()
//	validateStruct()
//}
//
//func validateStruct() {
//	user := User{
//		Accounts: []string{"account-1", "", "account-3"},
//		Balance: map[string]int{
//			"key1": 1,
//			"key2": -1,
//		},
//	}
//
//	err := validate.Struct(user)
//	if err != nil {
//		for _, err := range err.(validator.ValidationErrors) {
//			// Key: 'User.Accounts[1]' Error:Field validation for 'Accounts[1]' failed on the 'required' tag
//			// Key: 'User.Balance[key2]' Error:Field validation for 'Balance[key2]' failed on the 'gt' tag
//			fmt.Println(err.Error())
//		}
//	}
//}
