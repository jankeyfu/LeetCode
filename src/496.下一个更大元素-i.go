package src

/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 *
 * https://leetcode-cn.com/problems/next-greater-element-i/description/
 *
 * algorithms
 * Easy (62.25%)
 * Likes:    151
 * Dislikes: 0
 * Total Accepted:    20.7K
 * Total Submissions: 33K
 * Testcase Example:  '[4,1,2]\n[1,3,4,2]'
 *
 * 给定两个没有重复元素的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。找到 nums1 中每个元素在 nums2
 * 中的下一个比其大的值。
 *
 * nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出-1。
 *
 * 示例 1:
 *
 *
 * 输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
 * 输出: [-1,3,-1]
 * 解释:
 * ⁠   对于num1中的数字4，你无法在第二个数组中找到下一个更大的数字，因此输出 -1。
 * ⁠   对于num1中的数字1，第二个数组中数字1右边的下一个较大数字是 3。
 * ⁠   对于num1中的数字2，第二个数组中没有下一个更大的数字，因此输出 -1。
 *
 * 示例 2:
 *
 *
 * 输入: nums1 = [2,4], nums2 = [1,2,3,4].
 * 输出: [3,-1]
 * 解释:
 * 对于num1中的数字2，第二个数组中的下一个较大数字是3。
 * ⁠   对于num1中的数字4，第二个数组中没有下一个更大的数字，因此输出 -1。
 *
 *
 * 注意:
 *
 *
 * nums1和nums2中所有元素是唯一的。
 * nums1和nums2 的数组大小都不超过1000。
 *
 *
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := []int{}
	mp := map[int]int{}
	for i := 0; i < len(nums2); i++ {
		if len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
				mp[stack[len(stack)-1]] = nums2[i]
				stack = stack[:len(stack)-1]
			}
		}
		stack = append(stack, nums2[i])
	}
	ret := []int{}
	for _, num := range nums1 {
		if v, ok := mp[num]; ok {
			ret = append(ret, v)
		} else {
			ret = append(ret, -1)
		}
	}
	return ret
}

// @lc code=end
