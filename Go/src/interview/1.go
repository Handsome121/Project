/*
去重
*/
package main

import "fmt"

func main() {
	oldStr := []string{"1", "2", "3", "2", "2", "1"}
	fmt.Println("oldStr：", oldStr)
	mapString := make(map[string]bool)

	for _, num := range oldStr {
		mapString[num] = true
	}

	var newStr []string
	for key, _ := range mapString {
		newStr = append(newStr, key)
	}
	fmt.Println(newStr)
}
