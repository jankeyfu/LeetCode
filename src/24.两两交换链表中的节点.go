package src

/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	tmp := head
	for tmp != nil {
		if tmp.Next == nil {
			break
		}
		tmp.Val, tmp.Next.Val = tmp.Next.Val, tmp.Val
		tmp = tmp.Next.Next
	}
	return head
}

// @lc code=end
