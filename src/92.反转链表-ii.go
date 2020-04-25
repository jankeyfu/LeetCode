package src

/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 *
 * https://leetcode-cn.com/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (48.15%)
 * Likes:    263
 * Dislikes: 0
 * Total Accepted:    27.9K
 * Total Submissions: 57.9K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
 *
 * 说明:
 * 1 ≤ m ≤ n ≤ 链表长度。
 *
 * 示例:
 *
 * 输入: 1->2->3->4->5->NULL, m = 2, n = 4
 * 输出: 1->4->3->2->5->NULL
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	head = &ListNode{
		Next: head,
	}
	tmpH := head
	tmp := tmpH.Next
	i := 1
	for tmp.Next != nil {
		if i < m {
			tmpH = tmpH.Next
			tmp = tmpH.Next
		} else if i < n {
			nnext := tmp.Next.Next
			next := tmp.Next
			next.Next = tmpH.Next
			tmpH.Next = next
			tmp.Next = nnext
		} else {
			break
		}
		i++
	}

	return head.Next
}

// @lc code=end
