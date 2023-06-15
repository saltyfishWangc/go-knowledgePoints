package goroutine_pool

import (
	"sync"
	"sync/atomic"
	"testing"
)

var createNum int32

func createBuffer() interface{} {
	atomic.AddInt32(&createNum, 1)
	buffer := make([]byte, 1024)
	return buffer
}

// TestSyncPool
// sync.Pool除了最常见的池化提升性能的思路，最重要的是减少GC
// 常用于一些对象实例创建昂贵的场景。注意：Pool是Goroutine并发安全的

// sync.Pool适应场景
// sync.Pool本质用途是增加【临时对象】的重用率，减少GC负担
// sync.Pool中保存的元素有如下特征：
//
//	Pool池里的元素随时可能释放掉，释放策略完全由runtime内部管理；
//	Get获取到的元素可能是刚创建的，也可能是之前创建好cache的，使用者无法区分；
//	Pool池里面的元素个数你无法知道；

// 注意：临时对象。像socket这种带状态的、长期有效的资源是不适合Pool的。

// sync.Pool使用方法
//
//	1.初始化Pool实例New指定方法
//	2.申请对象Get
//	Get方法会返回Pool已经存在的对象，如果没有就使用New方法创建
//	3.释放对象Put
//	对象或资源不用时，调用Put方法把对象或资源放回池子，池子里面的对象啥时候真正释放是由go_routine进行回收，是不受外部控制的
func TestSyncPool(t *testing.T) {
	bufferPool := &sync.Pool{
		New: createBuffer,
	}

	workerPool := 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(workerPool)

	for i := 0; i < workerPool; i++ {
		go func() {
			defer wg.Done()
			buffer := bufferPool.Get()
			_ = buffer.([]byte)
			defer bufferPool.Put(buffer)
		}()
	}
	wg.Wait()
	// 如果不用sync.Pool，因为重复workPool次，所以会创建1024*1024个大小为1024的[]byte，用了sync.Pool管理后，只创建了createNum个
	t.Logf("%d buffer objects were create. \n", createNum)

}
