package main

import (
	"fmt"
	"time"
)

//时间类型
func timeDemo() {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

//时间戳:时间戳是自1970年1月1日（08:00:00GMT）至当前时间的总毫秒数。它也被称为Unix时间戳（UnixTimestamp）
func timestampDemo() {
	now := time.Now()                  //获取当前时间
	timestamp1 := now.Unix()           //时间戳
	timestamp2 := now.UnixNano()       //纳秒时间戳
	timestamp3 := now.UnixNano() / 1e6 //毫秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
	fmt.Printf("current timestamp2:%v\n", timestamp3)
}

//使用time.Unix()函数可以将时间戳转为时间格式。
func timestampDemo2(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

//时间间隔
//time.Duration是time包定义的一个类型，它代表两个时间点之间经过的时间，以纳秒为单位。
//time.Duration表示一段时间间隔，可表示的最长时间段大约290年。
//time.Duration表示1纳秒，time.Second表示1秒。

func Add() {
	now := time.Now()
	later := now.Add(time.Hour) // 当前时间加1小时后的时间
	fmt.Println(later)
}

func Sub() {
	now := time.Now()
	before := now.Add(-time.Hour)
	duration := now.Sub(before)
	fmt.Println(duration)
}

//Equal
//Before
//After

func formatDemo() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

func ParseTime() {
	now := time.Now()
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))
}

func main() {
	//timeDemo()
	//timestampDemo()
	//timestampDemo2(time.Now().Unix())
	//Add()
	//Sub()
	//formatDemo()
	ParseTime()

	//now := "2022-01-06T12:18:59Z"
	//timeFormat := "2006-01-02T15:04:05Z"
	//times, _ := time.Parse(timeFormat, now)
	//fmt.Println(times.Unix())
}
