package main

import (
	"sync"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(3)
		go func() {
			time.Sleep(2 * time.Second)
			t.Log("执行完了")
			wg.Done()
			//wg.Done()
			//wg.Done()
		}()
	}
	wg.Wait()
}
