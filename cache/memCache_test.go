package cache

import (
	"testing"
	"time"
)

func TestCacheDp(t *testing.T) {
	testData := []struct {
		key    string
		val    interface{}
		expire time.Duration
	}{
		{"slsd", 678, time.Second * 10},
		{"dds", 678, time.Second * 11},
		{"slsddsd", 678, time.Second * 12},
		{"dsd", map[string]interface{}{"a": 3, "b": false}, time.Second * 13},
		{"ds", "aadasdasdad", time.Second * 14},
		{"dsdsd", "这是什么", time.Second * 15},
	}

	c := NewMemCache()
	c.SetMaxMemory("10MB")
	for _, item := range testData {
		c.Set(item.key, item.val, item.expire)
		val, ok := c.Get(item.key)
		if !ok {
			t.Error("缓存取值失败")
		}
		t.Logf("获取缓存数据：%v", val)
	}
}
