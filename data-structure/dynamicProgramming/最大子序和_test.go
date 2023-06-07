package dynamicProgramming

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	t.Log(maxSubArray(nums))
}

// maxSubArray 最大子序和
// 给定一个整数数组nums，找出一个具有最大和的连续子数组(子数组最少包含一个元素)，返回其最大和
// 子数组是数组中的一个连续部分
func maxSubArray(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	// 定义dp数组。dp[i]表示数组中下标i为右端点的连续子数组的最大和
	dp := make([]int, len(nums))
	// 边界条件判断，当i等于0的时候，也就是前1个元素，能构成的最大和也就是它自己
	dp[0] = nums[0]
	maxC := dp[0]

	for i := 1; i < len(nums); i++ {
		// 要计算下标i为右端点的连续子数组的最大和，也就是计算dp[i]，只需要判断dp[i-1]是大于0还是小于0。如果dp[i-1]大于0，就继续累加，dp[i] = dp[i-1] + nums[i]
		// 如果dp[i-1]小于0，就需要直接把前面的舍弃，也就是说重新开始计算，否则会越加越小，直接让dp[i]=num[i]
		dp[i] = max(dp[i-1], 0) + nums[i]
		// 下标i为端点的连续子数组的最大和 和 最大和比较，取较大的
		maxC = max(maxC, dp[i])
	}
	return maxC
}
