package practise

import (
	"fmt"
	"sync"
	"testing"
)

// TestSwapPrintNum
// 两个协程用一个channel交替打印123456
func TestSwapPrintNum(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan struct{})
	go PrintOddNumber(&wg, ch, 6)
	go PrintEvenNumber(&wg, ch, 6)

	wg.Wait()
}

// PrintOddNumber
// 打印奇数
func PrintOddNumber(wg *sync.WaitGroup, ch chan struct{}, num int) {
	defer wg.Done()
	for i := 1; i <= num; i++ {
		ch <- struct{}{}
		if i%2 != 0 {
			fmt.Println("奇数：", i)
		}
	}
}

// PrintEvenNumber
// 打印偶数
func PrintEvenNumber(wg *sync.WaitGroup, ch chan struct{}, num int) {
	defer wg.Done()
	for i := 1; i <= num; i++ {
		<-ch
		if i%2 == 0 {
			fmt.Println("偶数：", i)
		}
	}
}
