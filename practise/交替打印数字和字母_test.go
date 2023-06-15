package practise

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

// TestSwapPrintNumAndLetter 交替打印数字和字母
// 题目：
// 使用两个goroutine交替打印序列，一个goroutine打印数字，另外一个goroutine打印字母，最终效果如下：
// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
func TestSwapPrintNumAndLetterRandom(t *testing.T) {
	printNumCh := make(chan struct{})
	printLetterCh := make(chan struct{})

	go func() {
		for {
			<-printNumCh
			t.Log(rand.Intn(10))
			printLetterCh <- struct{}{}
		}
	}()

	go func() {
		letterArr := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
		for {
			<-printLetterCh
			t.Logf("%c", letterArr[rand.Intn(26)])
			printNumCh <- struct{}{}
		}
	}()
	printNumCh <- struct{}{}
	time.Sleep(10 * time.Second)
}

// TestSwapPrintNumAndLetterSequence 交替顺序打印数字、字母
func TestSwapPrintNumAndLetterSequence(t *testing.T) {
	printNumCh := make(chan struct{})
	printLetterCh := make(chan struct{})

	// 退出通道
	exitCh := make(chan struct{})

	go func() {
		i := 1
		for {
			<-printNumCh
			t.Log(i)
			i++
			printLetterCh <- struct{}{}
		}
	}()

	go func() {
		letterArr := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
		i := 0
		for {
			<-printLetterCh
			t.Logf("%c", letterArr[i])
			i++
			if i >= len(letterArr) {
				exitCh <- struct{}{}
				break
			}
			printNumCh <- struct{}{}
		}
	}()
	printNumCh <- struct{}{}
	<-exitCh
	t.Log("完成")
}

// TestSwapPrintNumAndLetterWaitGroup 使用WaitGroup控制交替顺序打印数字、字母
func TestSwapPrintNumAndLetterWaitGroup(t *testing.T) {
	printNumCh := make(chan struct{})
	printLetterCh := make(chan struct{})

	var wg sync.WaitGroup

	go func() {
		i := 1
		for {
			<-printNumCh
			t.Log(i)
			i++
			printLetterCh <- struct{}{}
		}
	}()

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		letterArr := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
		i := 0
		for {
			<-printLetterCh
			t.Logf("%c", letterArr[i])
			i++
			if i >= len(letterArr) {
				wg.Done()
				break
			}
			printNumCh <- struct{}{}
		}
	}(&wg)
	printNumCh <- struct{}{}
	wg.Wait()
	t.Log("完成")
}
