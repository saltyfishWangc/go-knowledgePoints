package array

import "testing"

func TestRemoveDuplicate(t *testing.T) {
	//nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	nums := []int{1, 1, 2}

	newNums, length := removeDuplicate(nums)
	t.Logf("删除排序数组中的重复项后的数组是：%v 长度是：%d", newNums, length)
}

// removeDuplicate 删除排序数组中的重复项
// 给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
//
// 考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
//
// 更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums的其余元素与 nums 的大小不重要。
// 返回 k
func removeDuplicate(nums []int) ([]int, int) {
	if len(nums) == 1 || len(nums) == 0 {
		return nums, len(nums)
	}
	// 第一个指针指向第一个元素
	firstIndex := 0
	for i := 1; i < len(nums); i++ {
		// 从数组中第二个数开始遍历，如果和前面指针的数相等，则往后移动遍历
		// 如果和前面指针的数不相等，则firstIndex加1
		if nums[firstIndex] != nums[i] {
			firstIndex += 1
			// 让firstIndex下标前的数组是不重复的
			nums[firstIndex] = nums[i]
		}
	}
	// firstIndex指向的不重复数组中的最后一个元素下标，返回长度要加1返回
	return nums[:firstIndex+1], firstIndex + 1
}
