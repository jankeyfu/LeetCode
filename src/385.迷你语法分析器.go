package src

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=385 lang=golang
 *
 * [385] 迷你语法分析器
 *
 * https://leetcode-cn.com/problems/mini-parser/description/
 *
 * algorithms
 * Medium (39.36%)
 * Likes:    27
 * Dislikes: 0
 * Total Accepted:    2.7K
 * Total Submissions: 6.9K
 * Testcase Example:  '"324"'
 *
 * 给定一个用字符串表示的整数的嵌套列表，实现一个解析它的语法分析器。
 *
 * 列表中的每个元素只可能是整数或整数嵌套列表
 *
 * 提示：你可以假定这些字符串都是格式良好的：
 *
 *
 * 字符串非空
 * 字符串不包含空格
 * 字符串只包含数字0-9, [, - ,, ]
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 给定 s = "324",
 *
 * 你应该返回一个 NestedInteger 对象，其中只包含整数值 324。
 *
 *
 *
 *
 * 示例 2：
 *
 *
 * 给定 s = "[123,[456,[789]]]",
 *
 * 返回一个 NestedInteger 对象包含一个有两个元素的嵌套列表：
 *
 * 1. 一个 integer 包含值 123
 * 2. 一个包含两个元素的嵌套列表：
 * ⁠   i.  一个 integer 包含值 456
 * ⁠   ii. 一个包含一个元素的嵌套列表
 * ⁠        a. 一个 integer 包含值 789
 *
 *
 *
 *
 */

// @lc code=start
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func deserialize(s string) *NestedInteger {
	// 非 '[' 开头的表示只有数字
	if s[0] != '[' {
		i, _ := strconv.ParseInt(s, 10, 64)
		ret := &NestedInteger{}
		ret.SetInteger(int(i))
		return ret
	}
	sign := 1
	num := 0
	isNum := false
	stack := []NestedInteger{}
	for _, c := range s {
		if c == '-' {
			sign = -1
		} else if c >= '0' && c <= '9' {
			num = int(c-'0') + num*10
			isNum = true
		} else if c == '[' {
			stack = append(stack, NestedInteger{})
		} else if c == ',' || c == ']' {
			if isNum {
				cur := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				ni := &NestedInteger{}
				ni.SetInteger(sign * num)
				cur.Add(*ni)
				stack = append(stack, cur)
			}
			num, sign, isNum = 0, 0, false
			if c == ']' && len(stack) > 1 {
				cur := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack[len(stack)-1].Add(cur)
			}
		}
	}

	// fmt.Println(stack)
	return &stack[0]
}

// @lc code=end

type NestedInteger struct {
	Value int
}

func (n *NestedInteger) SetInteger(value int) {
	n.Value = value
}

func (n *NestedInteger) Add(elem NestedInteger) {

}
