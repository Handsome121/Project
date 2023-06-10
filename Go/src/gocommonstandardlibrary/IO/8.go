package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开文件
	//概念说明:file的叫法
	file, err := os.Open("E:/Project/Go/123.txt")
	if err != nil {
		fmt.Println("open gofile error=", err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //读到换行符就结束
		if err == io.EOF {                  //表示文件末尾
			break
		}
		fmt.Println(str)
	}
	fmt.Println("文件读取结束")
}
