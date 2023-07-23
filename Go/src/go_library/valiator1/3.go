///*
//跨结构体认证
//*/
//package main
//
//import (
//	"fmt"
//
//	"github.com/go-playground/validator/v10"
//)
//
//type User struct {
//	Uid     string `validate:"required,eqcsfield=Account.PayUid"`
//	Account Account
//}
//
//type Account struct {
//	PayUid string `validate:"required"`
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
//	account := Account{
//		PayUid: "uid-1025",
//	}
//
//	user := User{
//		Uid:     "uid-1024",
//		Account: account,
//	}
//
//	err := validate.Struct(user)
//	if err != nil {
//		for _, err := range err.(validator.ValidationErrors) {
//			fmt.Println(err.Error()) // Key: 'User.Uid' Error:Field validation for 'Uid' failed on the 'eqcsfield' tag
//		}
//	}
//}
//
//
//
