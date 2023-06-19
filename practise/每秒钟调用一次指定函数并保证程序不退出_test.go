package practise

import (
	"fmt"
	"testing"
	"time"
)

func TestProc(t *testing.T) {
	go func() {
		// 1 在这里需要你写算法
		// 2 要求每秒钟调用一次proc
		// 3 要求程序不能退出
		for {
			timer := time.NewTicker(time.Second)
			select {
			case <-timer.C:
				go func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Println(err)
						}
					}()
					proc()
				}()
			}
		}
	}()
	select {}
}

func proc() {
	panic("ok")
}
