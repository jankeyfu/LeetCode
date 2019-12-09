package main

/*
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。

注意你不能在买入股票前卖出股票。

示例 1:

输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
示例 2:

输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock
*/
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	minPrice := prices[0]
	dp := make([]int, len(prices))
	dp[0] = 0
	for i := 1; i < len(prices); i++ {
		minPrice = min(prices[i], minPrice)
		dp[i] = max(dp[i-1], prices[i]-minPrice)
	}
	return dp[len(prices)-1]
}

/*
题解：
将每一天卖出股票的最大收益作为状态
最大收益  = max(前一天的收益，当天股价 - 之前最小股价)
dp[i] = max(dp[i-1], prices[i]-minPrice)
*/
