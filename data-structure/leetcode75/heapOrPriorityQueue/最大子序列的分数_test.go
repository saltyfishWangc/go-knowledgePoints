package heapOrPriorityQueue

// 给你两个下标从 0开始的整数数组nums1和nums2，两者长度都是n，再给你一个正整数k。你必须从nums1中选一个长度为 k的 子序列对应的下标。
//
// 对于选择的下标i0，i1，...，ik - 1，你的分数定义如下：
//
//nums1中下标对应元素求和，乘以nums2中下标对应元素的最小值。
//用公示表示：(nums1[i0] + nums1[i1] +...+ nums1[ik - 1]) * min(nums2[i0] , nums2[i1], ... ,nums2[ik - 1])。
//请你返回 最大可能的分数。
//
//一个数组的 子序列下标是集合{0, 1, ..., n-1}中删除若干元素得到的剩余集合，也可以不删除任何元素。
//
//示例 1：
//
//输入：nums1 = [1,3,3,2], nums2 = [2,1,3,4], k = 3
//输出：12
//解释：
//四个可能的子序列分数为：
//- 选择下标 0 ，1 和 2 ，得到分数 (1+3+3) * min(2,1,3) = 7 。
//- 选择下标 0 ，1 和 3 ，得到分数 (1+3+2) * min(2,1,4) = 6 。
//- 选择下标 0 ，2 和 3 ，得到分数 (1+3+2) * min(2,3,4) = 12 。
//- 选择下标 1 ，2 和 3 ，得到分数 (3+3+2) * min(1,3,4) = 8 。
//所以最大分数为 12 。
//示例 2：
//
//输入：nums1 = [4,2,3,1,1], nums2 = [7,5,10,9,6], k = 1
//输出：30
//解释：
//选择下标 2 最优：nums1[2] * nums2[2] = 3 * 10 = 30 是最大可能分数。

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	return 0
}
