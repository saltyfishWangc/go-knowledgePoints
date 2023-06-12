package arrayAndString

import "testing"

// 给定一个整数数组nums，返回数组answer，其中answer[i]等于nums中除nums[i]之外其余各元素的乘积。
// 请不要使用除法，且在O(n)时间复杂度完成
// 进阶：在O(1)的额外空间复杂度内完成(出于对空间复杂度分析的目的，输出数组不被视为额外空间)
// 示例1：
// 输入：nums = [1,2,3,4]
// 输出：[24,12,8,6]

// 示例2：
// 输入：nums = [-1,1,0,-3,3]
// 输出：[0,0,9,0,0]
func TestProductExceptSelf(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	t.Log(productExceptSelf(nums))
}

// 采用O(n)的时间复杂度，O(1)的空间复杂度
// 遍历两遍数组，分别是从头到尾、从尾到头
func productExceptSelf(nums []int) []int {
	length := len(nums)
	answer := make([]int, length)
	// 从头到尾遍历数组，answer[i]表示下标为i的元素左边元素的乘积
	answer[0] = 1
	for i := 1; i < length; i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}

	// R表示i左边元素的乘积
	R := 1
	for i := length - 2; i >= 0; i-- {
		R = R * nums[i+1]
		answer[i] = answer[i] * R
	}
	return answer
}
