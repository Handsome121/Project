package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

const (
	// 控制输出日志信息的细节，不能控制输出的顺序和格式。
	// 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	Ldate         = 1 << iota     // 日期：2009/01/23
	Ltime                         // 时间：01:23:23
	Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
	Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
	Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
	LUTC                          // 使用UTC时间
	LstdFlags     = Ldate | Ltime // 标准logger的初始值
)

func init() {
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile) //配置日志输出文件位置
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate) //配置Log
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[pprof]") //配置前缀
	log.Println("这是一条很普通的日志。")
	//	v := "很普通的"
	//	log.Printf("这是一条%s日志。\n", v)
	//	log.Fatalln("这是一条会触发fatal的日志。")
	//	log.Panicln("这是一条会触发panic的日志。")
	//}
}

func main() {
	testing.Init()
}