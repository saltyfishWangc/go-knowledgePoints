package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestClosedChannelInSelect(t *testing.T) {
	ch := make(chan struct{})

	wg := sync.WaitGroup{}

	wg.Add(1)
	close(ch)
	go func() {
		for {
			select {
			case _, ok := <-ch:
				fmt.Println("拿到值")
				if !ok {
					fmt.Println("但是通道关闭了")
				}
				wg.Done()
				return
			case <-time.After(3 * time.Second):
				fmt.Println("超时了")
				wg.Done()
				return
			}
		}
	}()
	//ch <- struct{}{}
	wg.Wait()
}
