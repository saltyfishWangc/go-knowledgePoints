package main

import (
	"sync"
	"testing"
)

func TestMutexDoubleUnlockCanCatchPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error("捕获到异常:", err)
		}
	}()
	var mutex sync.Mutex
	t.Log("begin lock...\n")
	mutex.Lock()
	t.Log("get lock...\n")
	t.Log("unlock lock...\n")
	mutex.Unlock()
	t.Log("lock is unlock...\n")
	t.Log("unlock lock again...\n")
	mutex.Unlock()
}
