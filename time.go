package main

import (
	"fmt"
	"time"
)

func mainTime() {
	now := time.Now() // 获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     // 小时
	minute := now.Minute() //分钟
	second := now.Second() // 秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	fmt.Printf("%d-%2d-%2d %2d:%2d:%2d\n", year, month, day, hour, minute, second)
	fmt.Printf("%d-%d-%d %d:%d:%d\n", year, month, day, hour, minute, second)

	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp1:%v\n", timestamp2)
}
