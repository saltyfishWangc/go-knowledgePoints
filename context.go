package main

import (
	"bytes"
	"context"
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func mainContext() {
	//d := time.Now().Add(50 * time.Millisecond)
	//d := time.Now().Add(2 * time.Second)
	//ctx, cancelFunc := context.WithDeadline(context.Background(), d)
	//defer cancelFunc()
	//
	//select {
	//case <-time.After(time.Second):
	//	fmt.Println("overslept")
	//case <-ctx.Done():
	//	fmt.Println(ctx.Err())
	//}

	ctxa, cancel := context.WithCancel(context.Background())
	go work(ctxa, "work1")

	ctxb, _ := context.WithTimeout(ctxa, time.Second*3)
	go work(ctxb, "work2")

	ctxc := context.WithValue(ctxb, "key", "custom value")
	go workWithValue(ctxc, "work3")

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(time.Second)

}

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, " get message to quit")
			return
		default:
			println(name, " is running")
			time.Sleep(time.Second)
		}
	}
}

func workWithValue(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			println(name, " get message to quit")
			return
		default:
			value := ctx.Value("key").(string)
			println(name, " is running with value", value)
			time.Sleep(time.Second)
		}
	}
}

func ContextStudy() {
	start := time.Now()
	fmt.Println("起始时间：", start.String())

	dur := start.Add(time.Second * 10)
	fmt.Println("截止时间：", dur.String())

	// cancelFunc在被调用时，context就被关闭了，对应监听context的子协程就会收到通知
	// 所以，一般的，如果需要在截止时间前截止，就可以主动调用cancelFunc()，否则，会等到截止时间到
	deadlineCtx, cancelFunc := context.WithDeadline(context.Background(), dur)

	// WithTimeout和WitnDeadline是一样的，只是传入的参数不同，表示多久后结束，最终还是会调用WithDeadline
	//deadlineCtx, cancelFunc := context.WithTimeout(context.Background(), 10 * time.Second)

	// WithCancel对比WithTimeout和WithDeadline不同的是，没有截止时间或超时时间，想要停止，只有手动调用cancelFunc
	//ctx, cancelFunc := context.WithCancel(context.Background())

	// WithValue是没有cancelFunc，这种context的用途是上下文传值的，要是存在cancelFunc，上文调用了，那下文就空指针了，不合理，所以设计者对于这种context就没有cancelFunc
	//value := context.WithValue(context.Background(), "name", "wangc")

	go func(ctx context.Context, cancelFunc context.CancelFunc) {
		time.Sleep(5 * time.Second)
		if deadline, ok := ctx.Deadline(); ok {
			fmt.Println("现在时间：", deadline.String())
		}
		cancelFunc()
	}(deadlineCtx, cancelFunc)

	<-deadlineCtx.Done()
	deadlineTime := time.Now()
	fmt.Println("当前时间：", deadlineTime.String())
}

func ContextWithValueStudy() {
	valueContext := context.WithValue(context.Background(), "name", "wangc")
	go func(ctx context.Context) {
		localContext := context.WithValue(ctx, "name", "小明")

		fmt.Println("当前上下文的值：", localContext.Value("name"))
		fmt.Println("父节点上下文的值：", ctx.Value("name"))

	}(valueContext)

	time.Sleep(10 * time.Second)
}

func GetGouroutineId() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	parseUint, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(err)
	}
	return parseUint
}
