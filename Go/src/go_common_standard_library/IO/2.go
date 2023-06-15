package main

func main() {
	//os.Open()函数能够打开一个文件，返回一个*File和一个err。对得到的文件实例调用close()方法能够关闭文件。

	/*打开和关闭文件*/
	// 只读方式打开当前目录下的main.go文件
	//file, err := os.Open("./main.go")
	//if err != nil {
	//	fmt.Println("open file failed!, err:", err)
	//	return
	//}
	//// 关闭文件
	//_ = file.Close()

    /*写文件*/
	// 新建文件
	//file, err := os.Create("./123.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer file.Close()
	//for i := 0; i < concurrency; i++ {
	//	file.WriteString("ab\n") //写字符串
	//	file.Write([]byte("cd\n")) //写byte切片
	//}

	/*读文件*/
	// 打开文件
	//file, err := os.Open("./234.txt")
	//if err != nil {
	//	fmt.Println("open file err :", err)
	//	return
	//}
	//defer file.Close()
	//// 定义接收文件读取的字节数组
	//var buf [128]byte
	//var content []byte
	//for {
	//	n, err := file.Read(buf[:])
	//	if err == io.EOF {
	//		// 读取结束
	//		break
	//	}
	//	if err != nil {
	//		fmt.Println("read file err ", err)
	//		return
	//	}
	//	content = append(content, buf[:n]...)
	//}
	//fmt.Println(string(content))

	/*[拷贝文件]*/
	// 打开源文件
	//srcFile, err := os.Open("./123.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//// 创建新文件
	//dstFile, err2 := os.Create("./abc2.txt")
	//if err2 != nil {
	//	fmt.Println(err2)
	//	return
	//}
	//// 缓冲读取
	//buf := make([]byte, 1024)
	//for {
	//	// 从源文件读数据
	//	n, err := srcFile.Read(buf)
	//	if err == io.EOF {
	//		fmt.Println("读取完毕")
	//		break
	//	}
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	//写出去
	//	dstFile.Write(buf[:n])
	//}
	//srcFile.Close()
	//dstFile.Close()
}
