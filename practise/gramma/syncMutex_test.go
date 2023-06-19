package gramma

import (
	"fmt"
	"sync"
	"testing"
)

var mu sync.Mutex
var chain string

// golang中的锁是不可重入锁
func TestSyncMutex(t *testing.T) {
	chain = "main"
	A()
	fmt.Println(chain)
}

func A() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> A"
	B()
}
func B() {
	chain = chain + " --> B"
	C()
}
func C() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> C"
}

type MyMutex struct {
	count int
	sync.Mutex
}

func TestMutex3(t *testing.T) {
	var mu MyMutex
	mu.Lock()
	// mu赋值给mu2，同时把锁的状态也复制过去了，所以mu2里面的锁是上锁的状态
	var mu2 = mu
	mu.count++
	mu.Unlock()
	// mu2的锁本身就是上锁的状态，那这里就会报死锁
	mu2.Lock()
	mu2.count++
	mu2.Unlock()
	fmt.Println(mu.count, mu2.count)
}
