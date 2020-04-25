package src

/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,1]
输出: 1
示例 2:

输入: [4,1,2,1,2]
输出: 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/single-number
*/

func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ret := nums[0]
	for i := 1; i < len(nums); i++ {
		ret = ret ^ nums[i]
	}
	return ret
}
