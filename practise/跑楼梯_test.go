package practise

import "testing"

// TestUpStairs
// 爬楼梯：每次有两种情况，爬一个台阶或者两个台阶
func TestUpStairs(t *testing.T) {
	t.Log(upStairsRecursion(3))
	t.Log(upStairsDp(3))
}

// 递归的方式
// 缺点：重复计算
func upStairsRecursion(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	return upStairsRecursion(n-1) + upStairsRecursion(n-2)
}

// 动态规划的方式
// 存储已经计算的值，直接拿出来用
func upStairsDp(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
