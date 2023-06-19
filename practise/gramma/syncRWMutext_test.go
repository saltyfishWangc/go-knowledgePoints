package gramma

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var mu1 sync.RWMutex
var count int

// 还是死锁
func TestSyncRWMutex(t *testing.T) {
	go A1()
	time.Sleep(2 * time.Second)
	mu1.Lock()
	defer mu1.Unlock()
	count++
	fmt.Println(count)
}
func A1() {
	mu1.RLock()
	defer mu1.RUnlock()
	B1()
}
func B1() {
	time.Sleep(5 * time.Second)
	C1()
}
func C1() {
	mu1.RLock()
	defer mu1.RUnlock()
}
