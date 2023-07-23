///*
//自定义验证器
//*/
//package main
//
//import (
//	"database/sql"
//	"database/sql/driver"
//	"fmt"
//	"reflect"
//
//	"github.com/go-playground/validator/v10"
//)
//
//type DbBackedUser struct {
//	Name sql.NullString `validate:"required"`
//	Age  sql.NullInt64  `validate:"required"`
//}
//
//var validate *validator.Validate
//
//func main() {
//
//	validate = validator.New()
//
//	// 注册所有的 sql.Null* 类型，使用 ValidateValuer 自定义类型函数进行验证
//	validate.RegisterCustomTypeFunc(ValidateValuer, sql.NullString{}, sql.NullInt64{}, sql.NullBool{}, sql.NullFloat64{})
//
//	x := DbBackedUser{
//		Name: sql.NullString{String: "", Valid: true},
//		Age:  sql.NullInt64{Int64: 0, Valid: false},
//	}
//
//	err := validate.Struct(x)
//
//	if err != nil {
//		fmt.Printf("Err(s):\n%+v\n", err)
//	}
//}
//
//// ValidateValuer 实现了 validator.CustomTypeFunc 接口
//func ValidateValuer(field reflect.Value) interface{} {
//	// 如果字段的类型实现了 driver.Valuer 接口（即支持向数据库写入值），则进行验证
//	if valuer, ok := field.Interface().(driver.Valuer); ok {
//
//		val, err := valuer.Value() // 获取字段的值
//		if err == nil {
//			return val
//		}
//		// 处理错误
//	}
//
//	return nil
//}
//
//
//
