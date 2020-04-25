package src

/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
 *
 * https://leetcode-cn.com/problems/binary-tree-right-side-view/description/
 *
 * algorithms
 * Medium (64.00%)
 * Likes:    213
 * Dislikes: 0
 * Total Accepted:    40.5K
 * Total Submissions: 63.8K
 * Testcase Example:  '[1,2,3,null,5,null,4]'
 *
 * 给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
 *
 * 示例:
 *
 * 输入: [1,2,3,null,5,null,4]
 * 输出: [1, 3, 4]
 * 解释:
 *
 * ⁠  1            <---
 * ⁠/   \
 * 2     3         <---
 * ⁠\     \
 * ⁠ 5     4       <---
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	nQueue, lQueue, ret := []*TreeNode{}, []int{}, []int{}
	nQueue = append(nQueue, root)
	lQueue = append(lQueue, 1)
	curlevel, lastLevel, curNode, lastNode := 0, 0, &TreeNode{}, &TreeNode{}
	for len(nQueue) != 0 {
		curNode = nQueue[0]
		curlevel = lQueue[0]
		if curlevel != lastLevel && curlevel != 1 {
			ret = append(ret, lastNode.Val)
		}
		lastNode = curNode
		lastLevel = curlevel
		if len(nQueue) == 1 {
			nQueue = []*TreeNode{}
			lQueue = []int{}
		} else {
			nQueue = nQueue[1:]
			lQueue = lQueue[1:]
		}
		if curNode.Left != nil {
			nQueue = append(nQueue, curNode.Left)
			lQueue = append(lQueue, curlevel+1)
		}
		if curNode.Right != nil {
			nQueue = append(nQueue, curNode.Right)
			lQueue = append(lQueue, curlevel+1)
		}
	}
	ret = append(ret, lastNode.Val)
	return ret
}

// @lc code=end
