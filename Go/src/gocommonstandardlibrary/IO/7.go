package main

import (
	"encoding/csv"
	"os"
)

var data = [][]string{{"tom", "18", "beijing"}, {"jon", "19", "shanghai"}}

func main() {
	file, _ := os.Create("1.csv")
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	var header []string = []string{"name", "age", "address"} // 写入CSV文件头w.
	_ = writer.Write(header)
	// 方式以，逐条写入
	for _, value := range data {
		_ = writer.Write(value)
	}
	//_ = writer.Write([]string{"tom", "18", "beijing"})
	// 方式二，一次性写入，本质还是调用的Write(record []string)方法
	//_ = writer.WriteAll(data)
}

//func main() {
//	fmt.Println(data)
//}
