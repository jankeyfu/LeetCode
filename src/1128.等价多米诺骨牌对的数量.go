package src

/*
 * @lc app=leetcode.cn id=1128 lang=golang
 *
 * [1128] 等价多米诺骨牌对的数量
 *
 * https://leetcode-cn.com/problems/number-of-equivalent-domino-pairs/description/
 *
 * algorithms
 * Easy (39.00%)
 * Likes:    13
 * Dislikes: 0
 * Total Accepted:    3.9K
 * Total Submissions: 9.4K
 * Testcase Example:  '[[1,2],[2,1],[3,4],[5,6]]'
 *
 * 给你一个由一些多米诺骨牌组成的列表 dominoes。
 *
 * 如果其中某一张多米诺骨牌可以通过旋转 0 度或 180 度得到另一张多米诺骨牌，我们就认为这两张牌是等价的。
 *
 * 形式上，dominoes[i] = [a, b] 和 dominoes[j] = [c, d] 等价的前提是 a==c 且 b==d，或是 a==d 且
 * b==c。
 *
 * 在 0 <= i < j < dominoes.length 的前提下，找出满足 dominoes[i] 和 dominoes[j] 等价的骨牌对
 * (i, j) 的数量。
 *
 *
 *
 * 示例：
 *
 * 输入：dominoes = [[1,2],[2,1],[3,4],[5,6]]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= dominoes.length <= 40000
 * 1 <= dominoes[i][j] <= 9
 *
 *
 */

// @lc code=start
func numEquivDominoPairs(dominoes [][]int) int {
	pairMp := map[int]int{}
	var min, max int
	for _, dominoe := range dominoes {
		if dominoe[0] < dominoe[1] {
			min, max = dominoe[0], dominoe[1]
		} else {
			min, max = dominoe[1], dominoe[0]
		}
		add := min + max*10
		if v, ok := pairMp[add]; ok {
			pairMp[add] = v + 1
		} else {
			pairMp[add] = 1
		}
	}
	n := 0
	for _, v := range pairMp {
		n += v * (v - 1) / 2
	}
	return n
}

// @lc code=end

// 暴力法超时：
// for i := 0; i < len(dominoes)-1; i++ {
// 	for j := i + 1; j < len(dominoes); j++ {
// 		if (dominoes[i][0] == dominoes[j][0] && dominoes[i][1] == dominoes[j][1]) || (dominoes[i][0] == dominoes[j][1] && dominoes[i][1] == dominoes[j][0]) {
// 			n++
// 		}
// 	}
// }
