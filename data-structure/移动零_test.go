package data_structure

import "testing"

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	t.Log(moveZeroes(nums))
}

// moveZeroes
// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//请注意 ，必须在不复制数组的情况下原地对数组进行操作。

// 定义一个下标p指向数组中不存在零的序列的最后一位，遍历数组，如果遇到非0，就和p+1互换值，然后p+1。遇到零则不处理，直到将数组遍历完
func moveZeroes(nums []int) []int {
	// 指向数组中非零序列的最后一位
	p := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[p+1], nums[i] = nums[i], nums[p+1]
			p += 1
		}
	}
	return nums
}
