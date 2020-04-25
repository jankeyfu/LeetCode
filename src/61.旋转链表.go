package src

/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 *
 * https://leetcode-cn.com/problems/rotate-list/description/
 *
 * algorithms
 * Medium (39.47%)
 * Likes:    184
 * Dislikes: 0
 * Total Accepted:    35.7K
 * Total Submissions: 90.4K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
 *
 * 示例 1:
 *
 * 输入: 1->2->3->4->5->NULL, k = 2
 * 输出: 4->5->1->2->3->NULL
 * 解释:
 * 向右旋转 1 步: 5->1->2->3->4->NULL
 * 向右旋转 2 步: 4->5->1->2->3->NULL
 *
 *
 * 示例 2:
 *
 * 输入: 0->1->2->NULL, k = 4
 * 输出: 2->0->1->NULL
 * 解释:
 * 向右旋转 1 步: 2->0->1->NULL
 * 向右旋转 2 步: 1->2->0->NULL
 * 向右旋转 3 步: 0->1->2->NULL
 * 向右旋转 4 步: 2->0->1->NULL
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
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	length := 0
	for tmp := head; tmp != nil; tmp = tmp.Next {
		length++
	}
	firstP, secondP := head, head
	first, second := 1, 1
	for firstP.Next != nil {
		if first-second == k%length {
			second++
			secondP = secondP.Next
		}
		first++
		firstP = firstP.Next
	}
	// 尾结点连接头结点
	firstP.Next = head
	// 头结点
	head = secondP.Next
	// 断开尾结点
	secondP.Next = nil
	return head
}

// @lc code=end
