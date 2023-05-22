package main

import (
	"fmt"
	"testing"
	"time"
)

func TestSingleWayChannel(t *testing.T) {
	ch := make(chan int, 10)

	go func(c <-chan int) {
		for {
			fmt.Printf("读到数据：%d", <-ch)
			time.Sleep(1 * time.Second)
		}
	}(ch)

	go func(c chan<- int) {

	}(ch)
}

func TestChannel(t *testing.T) {
	ch := make(chan struct{})

	go func() {
		for {
			<-ch
			t.Log("协程A收到了")
		}
	}()

	go func() {
		for {
			<-ch
			t.Log("协程B收到了")
		}
	}()

	for {
		ch <- struct{}{}
		t.Log("往通道ch放数据")
		time.Sleep(1 * time.Second)
	}
}

// TestChannelDeadLock
// 当对于通道操作只有接收和发送只有一方时，编译时就会报死锁
func TestChannelDeadLock(t *testing.T) {
	ch := make(chan struct{})
	//<-ch
	ch <- struct{}{}
	t.Log("收到了")
}

// TestReadFromCloedChannel
// 从已关闭的通道读取数据：不管通道是否关闭，从通道中读取数据都不会报错，只是读出来的数据问题，如果通道已关闭，读取出来的数据是类型的零值
// 可以通道ok来判断通道是否已关闭
func TestReadFromCloedChannel(t *testing.T) {
	ch := make(chan struct{})
	//close(ch)

	go func() {
		for i := 0; i < 5; i++ {
			if _, ok := <-ch; ok {
				t.Log("从通道中读到值")
				close(ch)
			} else {
				t.Log("通道已关闭")
			}
		}
	}()

	ch <- struct{}{}

	time.Sleep(5 * time.Second)
}
