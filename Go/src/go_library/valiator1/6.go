/*
国际化翻译
*/

package main

import (
	"fmt"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

type User struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       uint8  `validate:"gte=0,lte=130"`
}

var validate *validator.Validate
var trans ut.Translator

func main() {
	validate = validator.New()

	// 中文翻译器
	uniTrans := ut.New(zh.New())
	trans, _ = uniTrans.GetTranslator("zh")
	// 注册翻译器到验证器
	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		_ = fmt.Sprintf("registerDefaultTranslations fail: %s\n", err.Error())
	}

	validateStruct()
}

func validateStruct() {
	user := &User{
		FirstName: "Badger",
		LastName:  "Smith",
		Age:       135,
	}

	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			// 翻译
			fmt.Println(err.Translate(trans)) // Age必须小于或等于130
		}
	}
}

