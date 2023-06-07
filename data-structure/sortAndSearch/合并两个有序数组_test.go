package sortAndSearch

import "testing"

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
}

// merge 合并两个有序数组
// 给定两个按非递减顺序排列的整数数组nums1和nums2，另有两个整数m和n，分别表示nums1和nums2中的元素数目
// 请合并nums2到nums1中，使合并后的数组同样按非递减顺序排列
// 注意：最终，合并后数组不应由函数返回，而是存储在数组nums1中。为了应对这种情况，nums1的初始长度为m+n，其中前m个元素表示应合并的元素，后n个元素为0。应忽略，nums2的长度为n
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 定义p指向nums1的尾部待插入的下标
	p := m + n
	i, j := m-1, n-1
	for ; i >= 0; i-- {
		for ; j >= 0; j-- {
			if nums1[i] > nums2[j] {
				nums1[p] = nums1[i]
				p--
				break
			} else {
				nums1[p] = nums2[j]
				p--
			}
		}
	}

	if j >= 0 {
		// 说明nums1的都遍历完了，nums2还没遍历完
		for ; j >= 0; j-- {
			nums1[p] = nums2[j]
			p--
		}
	}
}
