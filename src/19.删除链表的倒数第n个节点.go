package src

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	firstP, secondP := head, head
	first, second := 0, 0

	for firstP != nil {
		if first-second == n+1 {
			second++
			secondP = secondP.Next
		}
		first++
		firstP = firstP.Next
	}
	if first-second != n+1 {
		head = head.Next
	} else {
		secondP.Next = secondP.Next.Next
	}
	return head
}

// @lc code=end
