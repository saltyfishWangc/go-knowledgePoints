package data_structure

import "testing"

func TestReverse(t *testing.T) {
	t.Log(reverse(123))
}

func reverse(x int) int {
	// 正负标志位 false表示负数
	flag := true
	if x < 0 {
		flag = false
		// x是负数，先把x转成整数
		x = 0 - x
	}
	// 先将整数一次除以10，得到的余数加到数组中，这样就是反过来的
	res := 0
	for x != 0 {
		res = res*10 + x%10
		x = x / 10
	}
	if !flag {
		res = 0 - res
	}
	return res
}
