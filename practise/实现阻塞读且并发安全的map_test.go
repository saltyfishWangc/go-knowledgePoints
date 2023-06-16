package practise

import (
	"sync"
	"testing"
	"time"
)

// GO⾥⾯MAP如何实现key不存在 get操作等待 直到key存在或者超时，保证并发安全，
// 且需要实现以下接⼝：
//
//	type sp interface {
//		Out(key string, val interface{}) //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
//	    Rd(key string, timeout time.Duration) interface{} //读取一个key，如果key不存在阻塞，等待key存在或者超时
//	}
func TestConcurrentMap(t *testing.T) {

}

type sp interface {
	Out(key string, val interface{})
	Rd(key string, timeout time.Duration)
}

type Map struct {
	lock *sync.RWMutex
	c    map[string]*entry
}

type entry struct {
	val     interface{}
	ch      chan struct{}
	isExist bool
}

func (m Map) Out(key string, val interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	item, ok := m.c[key]
	if !ok {
		m.c[key] = &entry{
			val: val,
		}
		return
	}
	// 存在则覆盖值
	item.val = val
	// 判断是否有读的channel阻塞
	if !item.isExist {
		// 读协程那边没读到时，会新生成通道，然后一直在读通道阻塞在那，这边发现通道存在，关闭通道后，那边就收到了这边存入key的通知了，就可以读到值返回
		if item.ch != nil {
			close(item.ch)
			// 为什么这里关闭了通道，还要手动赋值nil，因为通道在关闭后并不是立刻就回收了，那item.ch的值并不就是nil，只是关闭状态的通道(关闭的通道并不就是为nil的通道)
			item.ch = nil
		}
	}
}

func (m Map) Rd(key string, timeout time.Duration) interface{} {
	m.lock.RLock()
	defer m.lock.RUnlock()
	item, ok := m.c[key]
	if !ok {
		ch := make(chan struct{})
		m.c[key] = &entry{
			isExist: false,
			ch:      ch,
		}
		for {
			select {
			// 那边通道关闭了，这边也会收到通知，只是值不可用，可以用v , ok := <- ch; ok 为false表示通道关闭
			case <-ch:
				i, _ := m.c[key]
				return i.val
			case <-time.After(timeout):
				return nil
			}
		}
	}
	return item.val
}
