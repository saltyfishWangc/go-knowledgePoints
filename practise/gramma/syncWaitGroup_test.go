package gramma

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// panic: sync: WaitGroup is reused before previous Wait has returned
func TestSyncWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		wg.Add(1)
	}()
	wg.Wait()
}

type Once struct {
	m    sync.Mutex
	done int32
}

func (o *Once) Do(f func()) {
	if o.done == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		o.done = 1
		f()
	}
}

func TestDoubleCheck(t *testing.T) {
	o := &Once{}
	o.Do(func() {
		fmt.Println("hello")
	})
}
