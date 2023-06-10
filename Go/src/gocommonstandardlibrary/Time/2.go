package main

import (
	"log"
	"time"
)

func AfterDemo() {
	log.Println(time.Now())
	<-time.After(1 * time.Second)
	//time.Sleep(1 * time.Second)
	log.Println(time.Now())
}

func AfterFuncDemo2() {
	log.Println("AfterFuncDemo start: ", time.Now())
	time.AfterFunc(5*time.Second, func() {
		log.Println("AfterFuncDemo end: ", time.Now())
	})
	time.Sleep(10 * time.Second)
}

func main() {
	//AfterDemo()
	AfterFuncDemo2()
}
