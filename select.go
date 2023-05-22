package main

import (
	"fmt"
	"time"
)

func mainSelect() {

	//output1 := make(chan string)
	//output2 := make(chan string)
	//
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	output1 <- "test1"
	//}()
	//
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	output2 <- "test2"
	//}()
	//
	////select {
	////case s1 := <-output1:
	////	fmt.Println("s1=", s1)
	////case s2 := <-output2:
	////	fmt.Println("s2=", s2)
	////}
	//for {
	//	select {
	//	case s1 := <-output1:
	//		fmt.Println("s1=", s1)
	//	case s2 := <-output2:
	//		fmt.Println("s2=", s2)
	//	}
	//}

	output3 := make(chan string, 10)

	go func(writer chan string) {
		for {
			select {
			case writer <- "hello":
				fmt.Println("write hello")
			default:
				fmt.Println("channel full")
			}
			time.Sleep(500 * time.Millisecond)
		}
	}(output3)

	for s := range output3 {
		fmt.Println("res:", s)
		time.Sleep(time.Second)
	}
}
