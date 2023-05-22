package main

import "fmt"

/*
*
3个函数分别打印cat、dog、fish，要求每个函数都要起一个goroutine，按照cat、dog、fish顺序打印在屏幕上100次
*/
func mainChannelPractise() {

	catChan := make(chan int)
	dogChan := make(chan int)
	fishChan := make(chan int)

	endChan := make(chan int)

	go func() {
		i := 0
		for {
			select {
			case <-fishChan:
				if i == 100 {
					endChan <- 1
					break
				} else {
					fmt.Println("cat")
					i++
					catChan <- 1
				}
			}
		}
	}()

	go func() {
		for _ = range catChan {
			fmt.Println("dog")
			dogChan <- 1
		}
	}()

	go func() {
		for _ = range dogChan {
			fmt.Println("fish")
			fishChan <- 1
		}
	}()

	fishChan <- 1

	<-endChan
}
