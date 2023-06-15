package sortAndSearch

import "testing"

func TestBubboSort(t *testing.T) {
	nums := []int{5, 6, 2, 1, 4}
	t.Log(bubboSort(nums))
}

// bubboSort 冒泡排序 时间复杂度O(n^2)
func bubboSort(nums []int) []int {
	length := len(nums)
	// 一共遍历length-1次
	for i := 0; i < length-1; i++ {
		// 每次从头开始比较当前和后面一个元素的大小，前者比后者大，就交换，这样一轮遍历后，最大的数就放到数组尾部
		for j := 0; j < length-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
