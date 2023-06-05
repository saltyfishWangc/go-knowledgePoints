package array

import "testing"

func TestSingleNumber(t *testing.T) {
	nums := []int{4, 1, 2, 1, 2}
	t.Log(singleNumber(nums))
}

// singleNumber 找出数组中唯一的数字
// 异或算法是0与任何数异或都是其本身
// 异或支持结合律
func singleNumber(nums []int) int {
	p := 0
	for _, num := range nums {
		p = p ^ num
	}
	return p
}
