package src

import "math"

/*
 * @lc app=leetcode.cn id=456 lang=golang
 *
 * [456] 132模式
 *
 * https://leetcode-cn.com/problems/132-pattern/description/
 *
 * algorithms
 * Medium (26.39%)
 * Likes:    147
 * Dislikes: 0
 * Total Accepted:    7.6K
 * Total Submissions: 28.9K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个整数序列：a1, a2, ..., an，一个132模式的子序列 ai, aj, ak 被定义为：当 i < j < k 时，ai < ak <
 * aj。设计一个算法，当给定有 n 个数字的序列时，验证这个序列中是否含有132模式的子序列。
 *
 * 注意：n 的值小于15000。
 *
 * 示例1:
 *
 *
 * 输入: [1, 2, 3, 4]
 *
 * 输出: False
 *
 * 解释: 序列中不存在132模式的子序列。
 *
 *
 * 示例 2:
 *
 *
 * 输入: [3, 1, 4, 2]
 *
 * 输出: True
 *
 * 解释: 序列中有 1 个132模式的子序列： [1, 4, 2].
 *
 *
 * 示例 3:
 *
 *
 * 输入: [-1, 3, 2, 0]
 *
 * 输出: True
 *
 * 解释: 序列中有 3 个132模式的的子序列: [-1, 3, 2], [-1, 3, 0] 和 [-1, 2, 0].
 *
 *
 */

// @lc code=start
func find132pattern(nums []int) bool {
	stack := []int{}
	ak := math.MinInt32
	for i := len(nums) - 1; i >= 0; i-- {
		if ak > nums[i] {
			return true
		}
		for len(stack) != 0 && stack[len(stack)-1] < nums[i] {
			ak = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}

// @lc code=end
