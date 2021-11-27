package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	filePath := "E:/Project/Go/123.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
	}

	defer file.Close()

	str := "hahahah\n"
	writer := bufio.NewWriter(file)

	for i := 0; i < 5; i++ {
		_, err := writer.WriteString(str)
		if err != nil {
			fmt.Printf("write string err=%v", err)
		}
	}
	_ = writer.Flush()
}
