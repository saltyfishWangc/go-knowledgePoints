package main

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

func TestContextStudy(t *testing.T) {
	ContextStudy()
}

func TestContextWithValueStudy(t *testing.T) {
	ContextWithValueStudy()
}

func TestContextScope(t *testing.T) {
	ctx := context.Background()
	fmt.Println("parent context:", &ctx)
	fmt.Println("parent gid:", GetGouroutineId())
	childCtx, cancelFunc := context.WithCancel(ctx)
	fmt.Println("child context:", &childCtx)
	cancelFunc()
	fmt.Println("chlid after cancel:", &childCtx)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("goroutine child gid:", GetGouroutineId())
		fmt.Println("goroutine child context:", &childCtx)
		newCtx := context.Background()
		fmt.Println("goroutine new context:", &newCtx)
	}()
	wg.Wait()
}
