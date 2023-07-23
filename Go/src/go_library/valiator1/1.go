///*
//基本使用
//*/
//package main
//
//import (
//	"fmt"
//	"github.com/go-playground/validator/v10"
//)
//
//type User struct {
//	FirstName string `validate:"required"`
//	LastName  string `validate:"required"`
//	Age       uint8  `validate:"gte=0,lte=130"`
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
//	user := &User{
//		FirstName: "Badger",
//		LastName:  "Smith",
//		Age:       135,
//	}
//
//	err := validate.Struct(user)
//	if err != nil {
//		for _, err := range err.(validator.ValidationErrors) {
//			fmt.Println(err.Namespace())       // User.Age
//			fmt.Println(err.Field())           // Age
//			fmt.Println(err.StructNamespace()) // User.Age
//			fmt.Println(err.StructField())     // Age
//			fmt.Println(err.Tag())             // lte
//			fmt.Println(err.ActualTag())       // lte
//			fmt.Println(err.Kind())            // uint8
//			fmt.Println(err.Type())            // uint8
//			fmt.Println(err.Value())           // 135
//			fmt.Println(err.Param())           // 130
//			fmt.Println(err.Error())           // Key: 'User.Age' Error:Field validation for 'Age' failed on the 'lte' tag
//			fmt.Println()
//		}
//	}
//}
