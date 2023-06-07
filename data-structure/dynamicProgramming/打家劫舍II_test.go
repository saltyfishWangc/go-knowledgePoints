package dynamicProgramming

// robII 打家劫舍II
// 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都围成一圈，这意味着第一个房屋和最后一个房屋是紧挨着的。
// 同时，相邻的房屋装有相互连通的防盗系统，如果相邻的房屋在同一晚上被小偷闯入，系统会自动报警
func robII(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([][]int, len(nums))
	dp[0][0] = 0
	dp[0][1] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}

	return max(dp[len(nums)-1][0], dp[len(nums)-1][1])
}
