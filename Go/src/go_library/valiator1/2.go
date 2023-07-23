//package main
//
//import (
//	"fmt"
//
//	"github.com/go-playground/validator/v10"
//)
//
//type User struct {
//	Password        string `validate:"required"`
//	ConfirmPassword string `validate:"required,eqfield=Password"`
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
//		Password:        "password",
//		ConfirmPassword: "pass",
//	}
//
//	err := validate.Struct(user)
//	if err != nil {
//		for _, err := range err.(validator.ValidationErrors) {
//			fmt.Println(err.Error()) // Key: 'User.ConfirmPassword' Error:Field validation for 'ConfirmPassword' failed on the 'eqfield' tag
//		}
//	}
//}
