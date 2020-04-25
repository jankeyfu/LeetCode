package src

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (40.21%)
 * Likes:    1348
 * Dislikes: 0
 * Total Accepted:    191.1K
 * Total Submissions: 470.3K
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 *
 *
 * 注意空字符串可被认为是有效字符串。
 *
 * 示例 1:
 *
 * 输入: "()"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "()[]{}"
 * 输出: true
 *
 *
 * 示例 3:
 *
 * 输入: "(]"
 * 输出: false
 *
 *
 * 示例 4:
 *
 * 输入: "([)]"
 * 输出: false
 *
 *
 * 示例 5:
 *
 * 输入: "{[]}"
 * 输出: true
 *
 */

// @lc code=start
func isValid(s string) bool {
	stack := []rune{}
	opMp := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else if len(stack) == 0 {
			return false
		} else {
			if opMp[stack[len(stack)-1]] == c {
				stack = stack[0 : len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

// @lc code=end
