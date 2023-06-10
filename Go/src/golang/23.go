package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{Title: "喜剧之王", Year: 1902, Price: 10, Actors: []string{"星爷", "zhangbozhi"}}
	//编码的过程，结构体---->json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error")
		return
	}
	fmt.Printf("jsonStr =%s\n", jsonStr)
	//解码的过程，jsonstr---->结构体
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error", err)
		return
	}
	fmt.Printf("%v\n", myMovie)
}
