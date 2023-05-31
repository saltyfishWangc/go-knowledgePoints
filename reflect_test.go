package main

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"
)

// TestDeepEqual reflect.DeepEqual是比较两个变量的值的，通过反射，如果字段和对应的值都相同，则相等
func TestDeepEqual(t *testing.T) {
	mapA := make(map[int]int, 3)
	mapA[1] = 1
	mapA[2] = 2
	mapA[3] = 3

	mapB := make(map[int]int, 3)
	mapB[1] = 1
	mapB[2] = 2
	mapB[3] = 3

	t.Logf("mapA的地址：%p mapB的地址：%p", mapA, mapB)
	t.Log("mapA是否与mapB相等：", reflect.DeepEqual(mapA, mapB))
	t.Log("mapA是否与mapB相等：", reflect.DeepEqual(&mapA, &mapB))
}

type Animal struct{}

func (a *Animal) Eat() {
	fmt.Println("吃东西")
}

// TestUseMethodNameCallMethod
// 反射通过方法名调用方法
func TestUseMethodNameCallMethod(t *testing.T) {
	animal := &Animal{}
	reflect.ValueOf(animal).MethodByName("Eat").Call([]reflect.Value{})
}

// TestPrintOrder
// 三个函数，交替打印cat、fish、dog，要求每个函数用一个goroutine，按照顺序打印100次
func TestPrintOrder(t *testing.T) {
	catCh := make(chan struct{})
	fishCh := make(chan struct{})
	dogCh := make(chan struct{})

	quit := make(chan struct{})

	go func() {
		for {
			<-dogCh
			t.Log("cat")
			catCh <- struct{}{}
		}
	}()

	go func() {
		for {
			<-catCh
			t.Log("fish")
			fishCh <- struct{}{}
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			<-fishCh
			t.Log("dog")
			dogCh <- struct{}{}
		}
		quit <- struct{}{}
	}()

	dogCh <- struct{}{}

	<-quit
}

// TestPrintOrder
// 三个函数，交替打印cat、fish、dog，要求每个函数用一个goroutine，按照顺序打印100次
func TestPrintOrder1(t *testing.T) {
	catCh := make(chan struct{})
	fishCh := make(chan struct{})
	dogCh := make(chan struct{})

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(3)
		go func() {
			<-dogCh
			t.Log("cat")
			catCh <- struct{}{}
			wg.Done()
		}()

		go func() {
			<-catCh
			t.Log("fish")
			fishCh <- struct{}{}
			wg.Done()
		}()

		go func() {
			<-fishCh
			t.Log("dog")
			wg.Done()
			dogCh <- struct{}{}
		}()
	}

	dogCh <- struct{}{}

	//time.Sleep(5 * time.Second)
	wg.Wait()
}

// TestSwapPrintLetterAndNumber 两个协程交替打印字母和数字
func TestSwapPrintLetterAndNumber(t *testing.T) {
	letterCh := make(chan struct{})
	numberCh := make(chan struct{})

	wg := sync.WaitGroup{}

	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			<-letterCh
			t.Log("打印字母：", "a")
			numberCh <- struct{}{}
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			<-numberCh
			t.Log("打印数字：", 1)
			if i == 9 {
				wg.Done()
			}
			letterCh <- struct{}{}
		}
		//wg.Done() 放在这里的话最后一次letterCh <- struct{}{}会阻塞，造成死锁，因为numberCh <- struct{}{}的线程早就不存在了
	}()

	letterCh <- struct{}{}
	//time.Sleep(5 * time.Second)
	wg.Wait()
}

// TestContextTimeoutCancel
// 启动2个goroutine 2秒后取消，第一个协程1秒执行完，第二个协程3秒执行完
func TestContextTimeoutCancel(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)

	go func(ctx1 context.Context) {
		time.Sleep(1 * time.Second)
		t.Log("第一个协程执行完")
	}(ctx)

	go func(ctx2 context.Context) {
		time.Sleep(3 * time.Second)
		t.Log("第二个协程执行完")
	}(ctx)

	<-ctx.Done()
}

// TestContextTimeoutCancel
// 启动2个goroutine 2秒后取消，第一个协程1秒执行完，第二个协程3秒执行完。协程执行完后通过channel通知，是否超时
func TestContextTimeoutCancel1(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)

	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	go func(ctx1 context.Context) {
		//after := time.After(time.Second)
		select {
		case <-time.After(time.Second):
			t.Log("第一个协程执行完毕")
			ch1 <- struct{}{}
		case <-ch2:
			t.Log("恭喜你啊第二个协程")
		case <-ctx.Done():
			t.Log("我是第一个协程，我超时了，对不起")
		}
	}(ctx)

	go func(ctx2 context.Context) {
		select {
		case <-time.After(time.Second * 3):
			t.Log("第二个协程执行完毕")
			ch2 <- struct{}{}
		case <-ch1:
			t.Log("恭喜你啊第一个协程")
		case <-ctx.Done():
			t.Log("我是第二个协程，我超时了，对不起")
		}
	}(ctx)

	<-ctx.Done()
}
