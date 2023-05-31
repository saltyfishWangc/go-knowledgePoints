package main

import (
	"runtime"
	"testing"
)

func TestGoMaxProcs(t *testing.T) {
	t.Log("最大可执行的goroutine数：", runtime.NumGoroutine())
	t.Log("cpu数量：", runtime.NumCPU())
	t.Log("系统架构：", runtime.GOARCH)
}
