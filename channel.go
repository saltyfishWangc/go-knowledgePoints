package main

import "fmt"

func mainChannel() {
	//c := make(chan int)
	//go func() {
	//	for i := 0; i < 5; i++ {
	//		fmt.Println("发送:%d", i)
	//		c <- i
	//	}
	//	close(c)
	//}()
	//
	//for {
	//	if data, ok := <-c; ok {
	//		fmt.Println("接收:%d", data)
	//	} else {
	//		break
	//	}
	//}
	//fmt.Println("main结束")

	//c := make(chan int, 10)
	//go func() {
	//	for i := 0; i < 5; i++ {
	//		c <- i
	//	}
	//	close(c)
	//}()
	////for data := range c {
	////	fmt.Println("从通道获取结果:", data)
	////}
	//for i := 0; i < 10; i++ {
	//	fmt.Println("从通道获取结果:", <-c)
	//}

	//ch := make(chan struct{})
	//go func() {
	//	defer close(ch)
	//	time.Sleep(5 * time.Second)
	//	//ch <- struct{}{}
	//}()
	//
	//fmt.Println("得到协程睡眠完")
	//fmt.Println("从通道ch中获取到：", <-ch)
	//fmt.Println("得到协程睡眠完了")

	// 注意这个例子：在for...range操作中,close(chan)这种事没有值可取的，因为通道关闭，for就退出了
	ch := make(chan struct{})
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			fmt.Println("往通道中放值")
			ch <- struct{}{}
		}
	}()

	for _ = range ch {
		fmt.Println("从通道中取值")
	}
}
