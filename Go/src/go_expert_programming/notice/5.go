package main

import (
	"encoding/json"
	"fmt"
)

// Person 结构体字段必须大写才能被json.marshal()解析
type Person struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
	Age  int32  `json:"age"`
}

func main() {
	person := Person{
		Name: "ma",
		Sex:  "male",
		Age:  10,
	}

	byteData, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("data is %v", string(byteData))
}
