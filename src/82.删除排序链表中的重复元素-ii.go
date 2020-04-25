package src

/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 *
 * https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * algorithms
 * Medium (44.62%)
 * Likes:    184
 * Dislikes: 0
 * Total Accepted:    26.2K
 * Total Submissions: 58.6K
 * Testcase Example:  '[1,2,3,3,4,4,5]'
 *
 * 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
 *
 * 示例 1:
 *
 * 输入: 1->2->3->3->4->4->5
 * 输出: 1->2->5
 *
 *
 * 示例 2:
 *
 * 输入: 1->1->1->2->3
 * 输出: 2->3
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
func deleteDuplicatesii(head *ListNode) *ListNode {
	head = &ListNode{
		Next: head,
	}
	flag := false
	tmpH := head
	for tmp := head; tmp.Next != nil; tmp = tmpH {
		for tmp.Next.Next != nil && tmp.Next.Next.Val == tmp.Next.Val {
			flag = true
			tmp.Next = tmp.Next.Next
		}
		if flag {
			flag = false
			tmpH.Next = tmpH.Next.Next
		} else {
			tmpH = tmpH.Next
		}
	}
	return head.Next
}

// @lc code=end
