package practise

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestPrintNumRandom(t *testing.T) {
	ch := make(chan int)

	end := make(chan struct{})

	go func() {
		for i := 0; i < 5; i++ {
			ch <- rand.Int()
		}
		close(ch)
	}()

	go func() {
		for {
			select {
			case v, ok := <-ch:
				if ok {
					t.Log(v)
				} else {
					fmt.Println("通道已关闭")
					end <- struct{}{}
					return
				}
			}
		}
	}()
	//go func() {
	//	for v := range ch {
	//		t.Log(v)
	//	}
	//	end <- struct{}{}
	//}()

	<-end
}
