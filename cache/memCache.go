package cache

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// memCache 缓存结构体
type memCache struct {
	// 最大内存
	maxMemorySize int64
	// 最大内存字符串表示
	maxMemorySizeStr string
	// 当前已使用内存大小
	currentMemorySize int64
	// 缓存键值对
	values map[string]*memCacheValue
	// 读写锁
	locker sync.RWMutex
	// 清除过期缓存时间间隔
	clearExpireItemTimeInterval time.Duration
}

// memCacheValue 缓存值结构体
type memCacheValue struct {
	// value值
	val interface{}
	// 过期时间
	expireTime time.Time
	// 有效时长
	expire time.Duration
	// value大小
	size int64
}

func NewMemCache() Cache {
	return &memCache{
		values:                      make(map[string]*memCacheValue),
		clearExpireItemTimeInterval: time.Second * 10,
	}
}

// SetMaxMemory 设置缓存最大内存
func (mc *memCache) SetMaxMemory(size string) bool {
	mc.maxMemorySize, mc.maxMemorySizeStr = ParseSize(size)
	return true
}

// Set 设置缓存
func (mc *memCache) Set(key string, val interface{}, expire time.Duration) bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()
	v := &memCacheValue{
		val:        val,
		expireTime: time.Now().Add(expire),
		expire:     expire,
		size:       GetValSize(val),
	}
	mc.del(key)
	mc.add(key, v)
	// 添加元素后，如果当前内存超过了最大内存，需要删除
	// TODO 后面可以考虑失效策略
	if mc.currentMemorySize > mc.maxMemorySize {
		mc.del(key)
		log.Println(fmt.Sprintf("max memory size %d", mc.maxMemorySize))
		return false
	}
	return true
}

// Get 根据key获取value
func (mc *memCache) Get(key string) (interface{}, bool) {
	mc.locker.RLock()
	defer mc.locker.RUnlock()
	value, ok := mc.get(key)
	if ok {
		// 判断缓存是否过期
		if value.expire != 0 && value.expireTime.Before(time.Now()) {
			// 过期则删除该缓存
			mc.del(key)
			return nil, false
		}
		return value.val, ok
	}
	return nil, false
}

// Del 删除key
func (mc *memCache) Del(key string) bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()
	mc.del(key)
	return true
}

func (mc *memCache) del(key string) {
	tmp, ok := mc.values[key]
	if ok && tmp != nil {
		mc.currentMemorySize -= tmp.size
		delete(mc.values, key)
	}
}

func (mc *memCache) add(key string, val *memCacheValue) {
	mc.values[key] = val
	mc.currentMemorySize += val.size
}

func (mc *memCache) get(key string) (*memCacheValue, bool) {
	value, ok := mc.values[key]
	return value, ok
}

// Exists 判断key是否存在
func (mc *memCache) Exists(key string) bool {
	mc.locker.RLock()
	defer mc.locker.RUnlock()
	_, ok := mc.get(key)
	return ok
}

// Clear 清空所有key
func (mc *memCache) Clear() bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()
	mc.values = make(map[string]*memCacheValue, 0)
	mc.currentMemorySize = 0
	return true
}

// Keys 获取缓存中所有key的数量
func (mc *memCache) Keys() int64 {
	mc.locker.RLock()
	defer mc.locker.RUnlock()
	return int64(len(mc.values))
}

// clearExpireItem 定期清除已过期的缓存
func (mc *memCache) clearExpireItem() {
	timeTicker := time.NewTicker(mc.clearExpireItemTimeInterval)
	defer timeTicker.Stop()
	for {
		select {
		case <-timeTicker.C:
			for key, item := range mc.values {
				if item.expire != 0 && time.Now().After(item.expireTime) {
					mc.locker.Lock()
					mc.del(key)
					mc.locker.Unlock()
				}
			}
		}
	}
}
