package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Post请求示例
func main() {
	url := "http://127.0.0.1:9090/post"
	// 表单数据
	//contentType := "application/x-www-form-urlencoded"
	//data := "name=枯藤&age=18"
	// json
	contentType := "application/json"
	data := `{"name":"枯藤","age":18}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed,err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
