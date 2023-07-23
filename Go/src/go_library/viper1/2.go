package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func ReadIni() {
	v := viper.New()
	v.AddConfigPath("./")   // 路径
	v.SetConfigName("2")   // 名称
	v.SetConfigType("ini") // 类型

	err := v.ReadInConfig() // 读取配置
	if err != nil {
		_ = fmt.Errorf("error is %v", err)
	}

	s := v.GetString("db.username")
	fmt.Printf("s: %v\n", s)

	i := v.GetInt("web.port")
	fmt.Printf("i: %v\n", i)
}

func main() {
	ReadIni()
}
