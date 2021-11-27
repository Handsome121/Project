package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	filePath := "E:/Project/Go/123.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str1, err := reader.ReadString('\n')
		if err == io.EOF {                  //表示文件末尾
			break
		}
		fmt.Print(str1)
	}
	str := "abd\r\n"
	writer := bufio.NewWriter(file)

	for i := 0; i < 5; i++ {
		_, err := writer.WriteString(str)
		if err != nil {
			fmt.Printf("write string err=%v", err)
		}
	}
	_ = writer.Flush()
}
