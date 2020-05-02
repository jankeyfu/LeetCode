package src

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个排序链表
 *
 * https://leetcode-cn.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (49.83%)
 * Likes:    627
 * Dislikes: 0
 * Total Accepted:    111.2K
 * Total Submissions: 217.7K
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * 合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
 *
 * 示例:
 *
 * 输入:
 * [
 * 1->4->5,
 * 1->3->4,
 * 2->6
 * ]
 * 输出: 1->1->2->3->4->4->5->6
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
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	ret := lists[0]
	for i := 1; i < len(lists); i++ {
		ret = mergeLists(ret, lists[i])
	}
	return ret
}

func mergeLists(node1, node2 *ListNode) *ListNode {
	cur := &ListNode{}
	root := cur
	tmp1, tmp2 := node1, node2
	for tmp1 != nil && tmp2 != nil {
		if tmp1.Val < tmp2.Val {
			cur.Next = tmp1
			cur = cur.Next
			tmp1 = tmp1.Next
		} else {
			cur.Next = tmp2
			cur = cur.Next
			tmp2 = tmp2.Next
		}
	}
	for tmp1 != nil {
		cur.Next = tmp1
		cur = cur.Next
		tmp1 = tmp1.Next
	}
	for tmp2 != nil {
		cur.Next = tmp2
		cur = cur.Next
		tmp2 = tmp2.Next
	}
	return root.Next
}

// @lc code=end
