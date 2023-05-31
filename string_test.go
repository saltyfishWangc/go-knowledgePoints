package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// TestString
// 字符串拼接的几种方式
func TestString(t *testing.T) {
	a1 := "nihao"
	a2 := "wangc"
	t.Log("使用+进行拼接：", a1+a2)

	t.Log("使用fmt.Sprintf()进行拼接：", fmt.Sprintf("%s%s", a1, "wangc"))

	a3 := []string{"nihao", "wangc"}
	t.Log("使用strings.Join()进行拼接", strings.Join(a3, " "))

	var bt bytes.Buffer
	bt.WriteString(a1)
	bt.WriteString(a2)
	t.Log("使用bytes.Buffer进行拼接：", bt.String())

	var bf strings.Builder
	bf.WriteString(a1)
	bf.WriteString(a2)
	t.Log("使用strings.Builder进行拼接：", bf.String())
}

// Test9Multi9 9*9乘法表
func Test9Multi9(t *testing.T) {
	res := "\n"
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			res = fmt.Sprintf("%s\t %d * %d = %d", res, j, i, i*j)
		}
		res = fmt.Sprintf("%s\n", res)
	}
	t.Log(res)
}
