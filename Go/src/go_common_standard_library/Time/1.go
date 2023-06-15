package main

import (
	"log"
	"time"
)

//使用场景1:设定超时时间
//func WaitChannel(conn <-chan string) bool {
//	timer := time.NewTimer(1 * time.Second)
//	select {
//	case <-conn:
//		timer.Stop()
//		return true
//	case <-timer.C:
//		println("WaitChannel timeout!")
//		return false
//	}
//}

//延迟执行某个方法
//func DelayFunction() {
//	timer := time.NewTimer(concurrency * time.Second)
//	select {
//	case <-timer.C:
//		log.Println("Delay 5s,start to do something")
//	}
//}

//等待指定时间
//func AfterDemo() {
//	log.Println(time.Now())
//	<-time.After(concurrency * time.Second)
//	log.Println(time.Now())
//}

//延迟方法调用
func AfterFuncDemo() {
	log.Println("AfterFuncDemo start: ", time.Now())
	time.AfterFunc(5*time.Second, func() {
		log.Println("AfterFuncDemo end: ", time.Now())
	})
	time.Sleep(10 * time.Second)
}
func main() {
	//AfterDemo()
	AfterFuncDemo()

}
