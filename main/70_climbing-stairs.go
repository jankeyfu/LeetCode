package main

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/climbing-stairs
*/
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp, dp1, dp2 := 0, 1, 2
	for i := 3; i <= n; i++ {
		dp = dp1 + dp2
		dp1 = dp2
		dp2 = dp
	}
	return dp
}

/*
题解：可以摸索规律
第一层 1种：f(1) = 1
第二层 2种：f(2) = 2 <=  1+1 || 2+0
第三层 3种：f(3) = f(1) + f(2) = 3 <= 第一层+2 第二层+1
以此类推得出 f(n) = f(n-1) + f(n-2)

可以使用斐波拉契数列做，也可以使用动态规划计算
*/