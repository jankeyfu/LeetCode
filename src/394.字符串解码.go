package src

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 *
 * https://leetcode-cn.com/problems/decode-string/description/
 *
 * algorithms
 * Medium (47.09%)
 * Likes:    188
 * Dislikes: 0
 * Total Accepted:    15.6K
 * Total Submissions: 32.3K
 * Testcase Example:  '"3[a]2[bc]"'
 *
 * 给定一个经过编码的字符串，返回它解码后的字符串。
 *
 * 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
 *
 * 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
 *
 * 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
 *
 * 示例:
 *
 *
 * s = "3[a]2[bc]", 返回 "aaabcbc".
 * s = "3[a2[c]]", 返回 "accaccacc".
 * s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".
 *
 *
 */

// @lc code=start
func decodeString(s string) string {
	stack := []string{}
	for _, c := range s {
		// fmt.Println(stack)
		if c != ']' {
			stack = append(stack, string(c))
			continue
		}
		// 出栈
		tmp := ""
		idx := len(stack) - 1
		// 取字符
		for stack[idx] != "[" {
			tmp = stack[idx] + tmp
			idx--
		}
		// 取数字
		nums := ""
		idx--
		for idx >= 0 && inSlice(stack[idx], []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}) {
			nums = stack[idx] + nums
			idx--
		}
		if idx < 0 {
			stack = []string{}
		} else {
			stack = stack[:idx+1]
		}
		num, _ := strconv.Atoi(nums)
		if num == 0 {
			num = 1
		}
		tmp = strings.Repeat(tmp, num)
		stack = append(stack, tmp)
	}
	ret := ""
	for _, s := range stack {
		ret += s
	}
	return ret
}
func inSlice(s string, ss []string) bool {
	for i := range ss {
		if ss[i] == s {
			return true
		}
	}
	return false
}

// @lc code=end
