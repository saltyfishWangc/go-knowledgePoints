package cache

import "time"

// Cache 定义缓存的接口
type Cache interface {
	// SetMaxMemory size : 1KB 100KB 1MB 2MB
	SetMaxMemory(size string) bool
	// Set 设置缓存
	Set(key string, val interface{}, expire time.Duration) bool
	// Get 根据key获取value
	Get(key string) (interface{}, bool)
	// Del 删除key
	Del(key string) bool
	// Exists 判断key是否存在
	Exists(key string) bool
	// Clear 清空所有key
	Clear() bool
	// Keys 获取缓存中所有key的数量
	Keys() int64
}
