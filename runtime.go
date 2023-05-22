package main

import (
	"fmt"
	"runtime"
)

func mainRuntime() {
	// runtime.Gosched()
	// 让出时间片，重新等待安排任务。注意：只会让一次
	//go func(s string) {
	//	for i := 0; i < 2; i++ {
	//		fmt.Println(s)
	//		time.Sleep(time.Second)
	//	}
	//}("worlld")
	//// 主协程
	//for i := 0; i < 2; i++ {
	//	// 切一下，再次分配任务
	//	runtime.Gosched()
	//	fmt.Println("hello")
	//}
	// runtime.Gosched() 结束

	// runtime.Goexit()
	// 退出当前协程，当前协程的代码不会再执行。当然，这是优雅退出，像defer这些该执行的还是会执行
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		// 因为在匿名函数中runtime.Goexit()当前协程退出了，所以这里的代码也不会再执行
		fmt.Println("A")
	}()
	for {
	}
	//runtime.Goexit() 结束
}
