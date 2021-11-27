package main

import (
	"fmt"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}
func main() {

	//判断文件或文件夹是否存在
	file1Path := "E:/Project/Go/123.txt"
	isExist, _ := PathExists(file1Path)
	fmt.Println(isExist)

}
