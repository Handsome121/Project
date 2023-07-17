/*
简单工厂模式
*/
package main

import "fmt"

//===============1.定义简单接口=================
type SimpleAPI interface {
	Say(content string) string
}

//===============2.实现对应的接口===============
//对象1
type Chinese struct{}

func (c *Chinese) Say(content string) string {
	fmt.Println("这里是Chinese对应的方法")
	return string("cn")
}

//对象2
type English struct{}

func (e *English) Say(content string) string {
	fmt.Println("这里是English对应的方法")
	return "en"
}

//===================3.调用对应的方法================
func NewAPI(apiName string) SimpleAPI {
	if apiName == "cn" {
		return &Chinese{}
	} else if apiName == "en" {
		return new(English)
	}
	return nil
}
func main() {
	newAPI1 := NewAPI("cn")
	res := newAPI1.Say("你好")
	fmt.Println("调用结果：", res)
	newAPI2 := NewAPI("en")
	res = newAPI2.Say("hello")
	fmt.Println("调用结果：", res)
}
