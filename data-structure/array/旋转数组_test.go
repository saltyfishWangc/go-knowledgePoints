package array

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	//nums := []int{1, 2, 3, 4, 5, 6, 7}
	//k := 3

	nums := []int{-1, -100, 3, 99}
	k := 2
	rotate(nums, k)
}

// 给定一个整数数组 nums，将数组中的元素向右轮转 k个位置，其中k是非负数。
//
// 示例 1:
//
// 输入: nums = [1,2,3,4,5,6,7], k = 3
// 输出: [5,6,7,1,2,3,4]
// 解释:
// 向右轮转 1 步: [7,1,2,3,4,5,6]
// 向右轮转 2 步: [6,7,1,2,3,4,5]
// 向右轮转 3 步: [5,6,7,1,2,3,4]
// 示例2:
//
// 输入：nums = [-1,-100,3,99], k = 2
// 输出：[3,99,-1,-100]
// 解释:
// 向右轮转 1 步: [99,-1,-100,3]
// 向右轮转 2 步: [3,99,-1,-100]
func rotate(nums []int, k int) {
	length := len(nums)

	// 要旋转k个数，那就循环k次，每次将整个数组的元素往后移动一次，最后一个元素变成第一个
	for i := 0; i < k; i++ {
		// 从最后一个元素开始，将元素和前一个元素互换，最后可以达到每个元素往后移一位，以及最后一个元素变成第一个元素的效果
		for j := length - 1; j > 0; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
	fmt.Println("旋转后数组：", nums)
}
