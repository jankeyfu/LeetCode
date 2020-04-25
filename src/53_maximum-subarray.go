package src

/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-subarray
*/
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	ret := dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1]+nums[i] <= nums[i] {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		ret = max(ret, dp[i])
	}
	return ret
}

/*
本题可以采用动态规划来求解：两数之和如果前一个数是负数，两数之和肯定比后一个数要小
dp[i] = max(dp[i-1]+ nums[i], nums[i])
而最大子序列之和一定是其中的一项，可以使用 ret 记录最大序列和
*/
