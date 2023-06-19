package practise

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// 场景：在⼀个⾼并发的web服务器中，要限制IP的频繁访问。现模拟100个IP同时并发访问服务器，每个IP要重复访问1000次。
// 每个IP三分钟之内只能访问⼀次。修改以下代码完成该过程，要求能成功输出 success:100
func TestMapReadAndWriteWithConcurrent(t *testing.T) {

}

type Ban struct {
	visitIPs map[string]time.Time
	lock     sync.Mutex
}

func NewBan(ctx context.Context) *Ban {
	ban := &Ban{visitIPs: make(map[string]time.Time)}
	go func() {
		timer := time.NewTimer(time.Minute)
		for {
			select {
			case <-timer.C:
				ban.lock.Lock()
				for k, v := range ban.visitIPs {
					if time.Now().Sub(v) >= time.Minute {
						delete(ban.visitIPs, k)
					}
				}
				ban.lock.Unlock()
				timer.Reset(time.Minute)
			case <-ctx.Done():
				return
			}
		}
	}()
	return ban
}
func (o *Ban) visit(ip string) bool {
	o.lock.Lock()
	defer o.lock.Unlock()
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	o.visitIPs[ip] = time.Now()
	return false
}

func Test1(t *testing.T) {
	success := int64(0)
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	wg := sync.WaitGroup{}
	ban := NewBan(ctx)
	wg.Add(1000 * 100)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func(j int) {
				defer wg.Done()
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					atomic.AddInt64(&success, 1)
				}
			}(j)
		}
	}
	wg.Wait()
	fmt.Println("success:", success)
}
