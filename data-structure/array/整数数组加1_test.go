package array

import "testing"

func TestPlusOne(t *testing.T) {
	digits := []int{4, 3, 2, 9}
	t.Log(plusOne(digits))
}

// plusOne 从最高元素开始加1，要是大于等于10，就继续往前追加，一旦小于10，就加一返回
func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 > 9 {
			digits[i] = 0
		} else {
			digits[i] = digits[i] + 1
			return digits
		}
	}
	return append([]int{1}, digits...)
}
