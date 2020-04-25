package src

/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 *
 * https://leetcode-cn.com/problems/reorder-list/description/
 *
 * algorithms
 * Medium (53.47%)
 * Likes:    138
 * Dislikes: 0
 * Total Accepted:    12.8K
 * Total Submissions: 23.8K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
 * 将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
 *
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 * 示例 1:
 *
 * 给定链表 1->2->3->4, 重新排列为 1->4->2->3.
 *
 * 示例 2:
 *
 * 给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
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
func reorderList(head *ListNode) {
	// 链表截成两段

	// 后半段翻转

	// 两段链表组合
}

// @lc code=end

func length(head *ListNode) int {
	len := 0
	for tmp := head; tmp != nil; tmp = tmp.Next {
		len++
	}
	return len
}

// func splitHalf(head *ListNode) []*ListNode {
// 获取链表长度
// len := length(head)
// if len <= 2 {
// 	return []*ListNode{head, nil}
// }
// i := 1
// for tmp := head; tmp != nil; tmp = tmp.Next {
// 	if i == len/2 {
// 		right = tmp.Next
// 	}
// }
// }
