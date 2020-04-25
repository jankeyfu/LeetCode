package src

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ret, root *ListNode

	step := 0
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + step
		if val >= 10 {
			step = 1
		} else {
			step = 0
		}
		val = val - 10*step
		if ret == nil {
			ret = &ListNode{
				Val: val,
			}
			root = ret
		} else {
			ret.Next = &ListNode{
				Val: val,
			}
			ret = ret.Next
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		val := l1.Val + step
		if val >= 10 {
			step = 1
		} else {
			step = 0
		}
		val = val - 10*step
		ret.Next = &ListNode{
			Val: val,
		}
		l1 = l1.Next
		ret = ret.Next
	}
	for l2 != nil {
		val := l2.Val + step
		if val >= 10 {
			step = 1
		} else {
			step = 0
		}
		val = val - 10*step
		ret.Next = &ListNode{
			Val: val,
		}
		l2 = l2.Next
		ret = ret.Next
	}
	if step > 0 {
		ret.Next = &ListNode{
			Val: step,
		}
	}
	return root
}

// @lc code=end
