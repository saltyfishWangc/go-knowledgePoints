package main

import (
	"fmt"
	"time"
)

func poc() {
	panic("出错了")
}

func RecoverTest() {
	// 每隔一秒调用poc方法，但是程序不能退出
	ticker := time.NewTicker(time.Second)
	go func() {
		select {
		case <-ticker.C:
			defer func() {
				if err := recover(); err != nil {
					fmt.Println(err)
				}
			}()
			poc()
		}
	}()

	select {}
}
