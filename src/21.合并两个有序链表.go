package src

import "fmt"

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var ret, head *ListNode
	for l1 != nil && l2 != nil {
		if ret == nil {
			if l1.Val <= l2.Val {
				ret = &ListNode{
					Val: l1.Val,
				}
				l1 = l1.Next
			} else {
				ret = &ListNode{
					Val: l2.Val,
				}
				l2 = l2.Next
			}
			head = ret
		} else {
			if l1.Val <= l2.Val {
				ret.Next = &ListNode{
					Val: l1.Val,
				}
				l1 = l1.Next
			} else {
				ret.Next = &ListNode{
					Val: l2.Val,
				}
				l2 = l2.Next
			}
			ret = ret.Next
		}
	}
	if l1 != nil {
		fmt.Println(l1)
		if ret == nil {
			ret = l1
			head = ret
		} else {
			ret.Next = l1
		}
	}
	if l2 != nil {
		if ret == nil {
			ret = l2
			head = ret
		} else {
			ret.Next = l2
		}
	}
	return head
}

// @lc code=end
