package src

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 *
 * https://leetcode-cn.com/problems/permutations/description/
 *
 * algorithms
 * Medium (74.70%)
 * Likes:    666
 * Dislikes: 0
 * Total Accepted:    118.8K
 * Total Submissions: 156.7K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个 没有重复 数字的序列，返回其所有可能的全排列。
 *
 * 示例:
 *
 * 输入: [1,2,3]
 * 输出:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 *
 */

// @lc code=start
func permute(nums []int) [][]int {
	return getOrder([]int{}, nums)
}

func getOrder(item, nums []int) [][]int {
	// fmt.Println(item, nums)
	ret := [][]int{}
	if len(nums) == 1 {
		item = append(item, nums[0])
		v := make([]int, len(item))
		copy(v, item)
		return [][]int{v}
	}
	for i := 0; i < len(nums); i++ {
		item = append(item, nums[i])
		tmp := []int{}
		if i == 0 {
			tmp = nums[1:]
		} else {
			tmp = append(tmp, nums[0:i]...)
			tmp = append(tmp, nums[i+1:]...)
			// fmt.Println(item, nums, tmp, i, nums[0:i-1], nums[i:])
		}
		// retbefore := ret
		v := getOrder(item, tmp)
		ret = append(ret, v...)
		item = item[:len(item)-1]
		// fmt.Println(retbefore, v, ret, item, nums, i)
	}
	return ret
}

// @lc code=end
