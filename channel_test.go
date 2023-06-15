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

type user struct {
	name string
	age  int8
}

func TestChannelCopy(t *testing.T) {
	u := user{name: "Ankur", age: 25}
	g := &u
	c := make(chan *user, 5)
	c <- g
	fmt.Println("当前：", g)

	g = &user{
		name: "Ankur Anand",
		age:  100,
	}
	go printUser(c, 2*time.Second)

	go modifyUser(g)
	time.Sleep(5 * time.Second)
	fmt.Println(g)

}

func TestChannelCopy1(t *testing.T) {
	u := user{name: "Ankur", age: 25}
	g := &u
	c := make(chan *user, 5)
	c <- g
	fmt.Println("当前：", g)

	//go printUser(c, 1*time.Second)
	go printUser(c, 3*time.Second)

	//time.Sleep(2 * time.Second)
	go modifyUser(g)

	time.Sleep(5 * time.Second)
	fmt.Println("最后：", g)

}

func modifyUser(u *user) {
	fmt.Println("modifyUser Received Value", u)
	u.name = "wangc"
}

func printUser(ch <-chan *user, dau time.Duration) {
	time.Sleep(dau)
	fmt.Printf("after sleep：%f printUser goRoutine called：%v\n", dau.Seconds(), <-ch)
}

// TestHappensBefore
// go语言中关于channel的happened-before原则，对于无缓冲和有缓冲的chan的不同：
// 非缓冲的chan中，第n个receive一定happended-before第n个send finished。注意这里说的是send finished的，并不是说send
// 下面的例子就说明了：我们在主协程中receive后直接打印，但是主协程并没有直接跑完，而是等着赋值的goroutine
// 如果将done改成缓冲的chan，那主协程中的打印就是空的
func TestHappensBefore(t *testing.T) {
	//var done = make(chan bool, 1)
	var done = make(chan bool)
	var msg string
	go func() {
		msg = "hello world"
		time.Sleep(2 * time.Second)
		<-done
	}()
	done <- true
	t.Log(msg)
}

func TestMultiReceive(t *testing.T) {
	ch := make(chan struct{})
	go func() {
		<-ch
		t.Log("A收到了")
	}()

	go func() {
		<-ch
		t.Log("B收到了")
	}()

	ch <- struct{}{}

	time.Sleep(2 * time.Second)
}

func TestChanPanic(t *testing.T) {
	//1 给一个nil channel发送数据，会造成死锁
	//var s chan int
	//s <- 1
	//t.Log("hello")

	//2 从一个nil channel读取数据，会永久堵塞(也就是死锁)
	//var s chan int
	//<-s
	//t.Log("hello")

	//3 给一个已经关闭的channel发送数据，会引起panic
	//s := make(chan int)
	//close(s)
	//s <- 1
	//t.Log("hello")

	//4 从一个已经关闭的channel读取数据，返回的是类型的零值
	s := make(chan int)
	close(s)
	t.Log("hello ", <-s)
}
