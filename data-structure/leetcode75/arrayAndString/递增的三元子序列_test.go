package arrayAndString

import "testing"

func TestIncreasingTriplet(t *testing.T) {
	nums := []int{3, 2, 5, 6, 4}
	increasingTriplet(nums)
}

func increasingTriplet(nums []int) bool {
	length := len(nums)
	if length < 3 {
		return false
	}

	// 表示nums[i]左边最小的数
	leftMin := make([]int, length)
	// 第一个元素左边没有数，所以是本身
	leftMin[0] = nums[0]
	for i := 1; i < length; i++ {
		leftMin[i] = min(leftMin[i-1], nums[i])
	}

	// 表示nums[i]右边最大的数
	rightMax := make([]int, length)
	rightMax[length-1] = nums[length-1]
	for i := length - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], nums[i])
	}

	for i := 1; i < length-1; i++ {
		if leftMin[i-1] < nums[i] && nums[i] < rightMax[i+1] {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
